package chore

import (
	"context"
	"fmt"
	"testing"
	"time"

	"donetick.com/core/config"
	chModel "donetick.com/core/internal/chore/model"
	cModel "donetick.com/core/internal/circle/model"
	syncModel "donetick.com/core/internal/sync/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newScoringTestRepository(t *testing.T) (*ChoreRepository, *gorm.DB) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(
		&chModel.Chore{}, &chModel.ChoreHistory{}, &chModel.ChoreAssignees{}, &chModel.TimeSession{},
		&cModel.UserCircle{}, &syncModel.SyncCursor{}, &syncModel.Tombstone{},
	))
	cfg := &config.Config{}
	cfg.Database.Type = "sqlite"
	return NewChoreRepository(db, cfg), db
}

func seedScoredChore(t *testing.T, db *gorm.DB, balance int) *chModel.Chore {
	t.Helper()
	due := time.Date(2026, 8, 26, 18, 0, 0, 0, time.UTC)
	points := 10
	assignee := 1
	chore := &chModel.Chore{
		ID: 10, Name: "Pay bill", FrequencyType: chModel.FrequencyTypeOnce,
		NextDueDate: &due, AssignedTo: &assignee, AssignStrategy: chModel.AssignmentStrategyKeepLastAssigned,
		IsActive: true, CircleID: 1, CreatedBy: 1, UpdatedBy: 1, Points: &points,
		TimingMode: chModel.TimingModeDeadline, SyncVersion: 1,
	}
	require.NoError(t, db.Create(&cModel.UserCircle{UserID: 1, CircleID: 1, IsActive: true, Points: balance}).Error)
	require.NoError(t, db.Create(&syncModel.SyncCursor{CircleID: 1, EntityType: syncModel.EntityTypeChore, MaxVersion: 1}).Error)
	require.NoError(t, db.Create(chore).Error)
	return chore
}

func TestCompleteChoreAppliesScoreExactlyOnce(t *testing.T) {
	repository, db := newScoringTestRepository(t)
	chore := seedScoredChore(t, db, 2)
	completedAt := chore.NextDueDate.Add(20 * time.Minute)
	score := &chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -2, Total: 8}

	require.NoError(t, repository.CompleteChore(context.Background(), chore, nil, 1, nil, &completedAt, nil, score))
	var member cModel.UserCircle
	require.NoError(t, db.Where("user_id = ? AND circle_id = ?", 1, 1).First(&member).Error)
	require.Equal(t, 10, member.Points)

	var history chModel.ChoreHistory
	require.NoError(t, db.Where("chore_id = ?", chore.ID).First(&history).Error)
	require.Equal(t, 8, *history.Points)
	require.Equal(t, 10, *history.BasePoints)
	require.Equal(t, -2, *history.TimingAdjustment)
	require.Equal(t, 0, *history.RecoveryPoints)

	err := repository.CompleteChore(context.Background(), chore, nil, 1, nil, &completedAt, nil, score)
	require.ErrorIs(t, err, ErrChoreChanged)
	var count int64
	require.NoError(t, db.Model(&chModel.ChoreHistory{}).Where("chore_id = ?", chore.ID).Count(&count).Error)
	require.EqualValues(t, 1, count)
	require.NoError(t, db.Where("user_id = ? AND circle_id = ?", 1, 1).First(&member).Error)
	require.Equal(t, 10, member.Points)
}

func TestApplyScoreNeverDropsCirclePointsBelowZero(t *testing.T) {
	repository, db := newScoringTestRepository(t)
	chore := seedScoredChore(t, db, 2)
	recordedAt := chore.NextDueDate.Add(9 * time.Hour)
	score := &chModel.ScoreResult{TimingAdjustment: -3, Total: -3}
	require.NoError(t, repository.RecordMissedScore(context.Background(), chore, recordedAt, score))
	var member cModel.UserCircle
	require.NoError(t, db.Where("user_id = ? AND circle_id = ?", 1, 1).First(&member).Error)
	require.Zero(t, member.Points)
}

func TestRecordMissedScoreIsIdempotentAndRecoveryRestoresCredit(t *testing.T) {
	repository, db := newScoringTestRepository(t)
	chore := seedScoredChore(t, db, 10)
	missedAt := chore.NextDueDate.Add(9 * time.Hour)
	missScore := &chModel.ScoreResult{TimingAdjustment: -3, Total: -3}
	require.NoError(t, repository.RecordMissedScore(context.Background(), chore, missedAt, missScore))
	require.ErrorIs(t, repository.RecordMissedScore(context.Background(), chore, missedAt.Add(time.Minute), missScore), ErrMissAlreadyScored)

	completedAt := missedAt.Add(time.Minute)
	recovery := &chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -7, RecoveryPoints: 3, Total: 6}
	require.NoError(t, repository.CompleteChore(context.Background(), chore, nil, 1, nil, &completedAt, nil, recovery))

	var member cModel.UserCircle
	require.NoError(t, db.Where("user_id = ? AND circle_id = ?", 1, 1).First(&member).Error)
	require.Equal(t, 13, member.Points) // 10 - 3 + 6
	var histories []chModel.ChoreHistory
	require.NoError(t, db.Order("id asc").Find(&histories).Error)
	require.Len(t, histories, 2)
	require.Equal(t, chModel.ChoreHistoryStatusMissed, histories[0].Status)
	require.Equal(t, chModel.ChoreHistoryStatusCompleted, histories[1].Status)
}

func TestApproveChoreUsesPendingHistoryPerformerAndScore(t *testing.T) {
	repository, db := newScoringTestRepository(t)
	chore := seedScoredChore(t, db, 0)
	chore.Status = chModel.ChoreStatusPendingApproval
	require.NoError(t, db.Model(&chModel.Chore{}).Where("id = ?", chore.ID).Update("status", chore.Status).Error)
	performed := chore.NextDueDate.Add(20 * time.Minute)
	require.NoError(t, db.Create(&chModel.ChoreHistory{
		ChoreID: chore.ID, PerformedAt: &performed, CompletedBy: 1, AssignedTo: chore.AssignedTo,
		DueDate: chore.NextDueDate, Status: chModel.ChoreHistoryStatusPendingApproval, SyncVersion: 1,
	}).Error)
	score := &chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -2, Total: 8}
	require.NoError(t, repository.ApproveChore(context.Background(), chore, 99, nil, nil, score))
	var member cModel.UserCircle
	require.NoError(t, db.Where("user_id = ? AND circle_id = ?", 1, 1).First(&member).Error)
	require.Equal(t, 8, member.Points)
}
