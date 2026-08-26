package chore

import (
	"math"
	"time"

	chModel "donetick.com/core/internal/chore/model"
)

const (
	scoringGracePeriod     = 15 * time.Minute
	scoringFirstLateStage  = 2 * time.Hour
	scoringSecondLateStage = 8 * time.Hour
)

type ScoreOutcome string

const (
	ScoreOutcomeCompleted ScoreOutcome = "completed"
	ScoreOutcomeMissed    ScoreOutcome = "missed"
	ScoreOutcomeRecovered ScoreOutcome = "recovered"
)

type ScoreInput struct {
	BasePoints  int
	TimingMode  chModel.TimingMode
	EarlyBonus  bool
	DueAt       *time.Time
	PerformedAt time.Time
	Outcome     ScoreOutcome
	Timezone    string
}

// ScoreTaskResult is the single pure scoring seam. It has no database,
// notification, or UI dependencies.
func ScoreTaskResult(input ScoreInput) chModel.ScoreResult {
	if input.BasePoints <= 0 {
		return chModel.ScoreResult{}
	}

	if input.Outcome == ScoreOutcomeMissed {
		penalty := percentage(input.BasePoints, 30)
		return chModel.ScoreResult{
			TimingAdjustment: -penalty,
			Total:            -penalty,
		}
	}

	earned := timedReward(input)
	result := chModel.ScoreResult{
		BasePoints:       input.BasePoints,
		TimingAdjustment: earned - input.BasePoints,
		Total:            earned,
	}
	if input.Outcome == ScoreOutcomeRecovered {
		result.RecoveryPoints = percentage(input.BasePoints, 30)
		result.Total += result.RecoveryPoints
	}
	return result
}

func timedReward(input ScoreInput) int {
	if input.TimingMode == "" || input.TimingMode == chModel.TimingModeUntimed || input.DueAt == nil {
		return input.BasePoints
	}

	performedAt := input.PerformedAt.UTC()
	dueAt := input.DueAt.UTC()

	if input.TimingMode == chModel.TimingModeToday {
		if !performedAt.After(endOfLocalDay(dueAt, input.Timezone)) {
			return input.BasePoints
		}
		return percentage(input.BasePoints, 30)
	}

	if performedAt.Before(dueAt) {
		if input.TimingMode == chModel.TimingModeDeadline && input.EarlyBonus {
			return input.BasePoints + percentage(input.BasePoints, 20)
		}
		return input.BasePoints
	}

	lateBy := performedAt.Sub(dueAt)
	switch {
	case lateBy <= scoringGracePeriod:
		return input.BasePoints
	case lateBy <= scoringFirstLateStage:
		return percentage(input.BasePoints, 80)
	case lateBy <= scoringSecondLateStage:
		return percentage(input.BasePoints, 60)
	default:
		return percentage(input.BasePoints, 30)
	}
}

func missBoundary(chore *chModel.Chore) *time.Time {
	if chore == nil || chore.NextDueDate == nil {
		return nil
	}
	var boundary time.Time
	switch chore.TimingMode {
	case chModel.TimingModeToday:
		boundary = endOfLocalDay(*chore.NextDueDate, choreTimezone(chore))
	case chModel.TimingModeDeadline, chModel.TimingModeWindow:
		boundary = chore.NextDueDate.UTC().Add(scoringSecondLateStage)
	default:
		return nil
	}
	return &boundary
}

func scoreForChore(chore *chModel.Chore, performedAt time.Time, outcome ScoreOutcome) *chModel.ScoreResult {
	if chore == nil || chore.Points == nil || *chore.Points <= 0 {
		return nil
	}
	result := ScoreTaskResult(ScoreInput{
		BasePoints:  *chore.Points,
		TimingMode:  chore.TimingMode,
		EarlyBonus:  chore.EarlyBonus,
		DueAt:       chore.NextDueDate,
		PerformedAt: performedAt,
		Outcome:     outcome,
		Timezone:    choreTimezone(chore),
	})
	return &result
}

func completionOutcome(histories []*chModel.ChoreHistory, dueAt *time.Time) ScoreOutcome {
	if dueAt == nil {
		return ScoreOutcomeCompleted
	}
	for _, history := range histories {
		if history == nil || history.Status != chModel.ChoreHistoryStatusMissed || history.DueDate == nil {
			continue
		}
		if history.DueDate.Equal(*dueAt) {
			return ScoreOutcomeRecovered
		}
	}
	return ScoreOutcomeCompleted
}

func choreTimezone(chore *chModel.Chore) string {
	if chore != nil && chore.FrequencyMetadataV2 != nil {
		return chore.FrequencyMetadataV2.Timezone
	}
	return ""
}

func endOfLocalDay(value time.Time, timezone string) time.Time {
	location := time.UTC
	if timezone != "" {
		if loaded, err := time.LoadLocation(timezone); err == nil {
			location = loaded
		}
	}
	local := value.In(location)
	return time.Date(local.Year(), local.Month(), local.Day(), 23, 59, 59, 999999999, location).UTC()
}

func percentage(value, percent int) int {
	return int(math.Round(float64(value*percent) / 100))
}
