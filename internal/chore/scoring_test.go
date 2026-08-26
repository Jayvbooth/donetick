package chore

import (
	"testing"
	"time"

	chModel "donetick.com/core/internal/chore/model"
	"github.com/stretchr/testify/require"
)

func TestScoreTaskResult(t *testing.T) {
	due := time.Date(2026, 8, 26, 18, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		input    ScoreInput
		expected chModel.ScoreResult
	}{
		{
			name:     "untimed keeps fixed points",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeUntimed, PerformedAt: due},
			expected: chModel.ScoreResult{BasePoints: 10, Total: 10},
		},
		{
			name:     "deadline useful early bonus is opt in",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, EarlyBonus: true, DueAt: &due, PerformedAt: due.Add(-time.Hour)},
			expected: chModel.ScoreResult{BasePoints: 10, TimingAdjustment: 2, Total: 12},
		},
		{
			name:     "window never receives early bonus",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeWindow, EarlyBonus: true, DueAt: &due, PerformedAt: due.Add(-time.Hour)},
			expected: chModel.ScoreResult{BasePoints: 10, Total: 10},
		},
		{
			name:     "grace period keeps full reward",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(15 * time.Minute)},
			expected: chModel.ScoreResult{BasePoints: 10, Total: 10},
		},
		{
			name:     "first late stage earns eighty percent",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(20 * time.Minute)},
			expected: chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -2, Total: 8},
		},
		{
			name:     "second late stage earns sixty percent",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(3 * time.Hour)},
			expected: chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -4, Total: 6},
		},
		{
			name:     "late floor earns thirty percent",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(9 * time.Hour)},
			expected: chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -7, Total: 3},
		},
		{
			name:     "today mode uses the stored local day",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeToday, DueAt: &due, PerformedAt: time.Date(2026, 8, 27, 3, 30, 0, 0, time.UTC), Timezone: "America/New_York"},
			expected: chModel.ScoreResult{BasePoints: 10, Total: 10},
		},
		{
			name:     "miss applies one bounded penalty",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(9 * time.Hour), Outcome: ScoreOutcomeMissed},
			expected: chModel.ScoreResult{TimingAdjustment: -3, Total: -3},
		},
		{
			name:     "recovery restores the missed deduction plus late floor",
			input:    ScoreInput{BasePoints: 10, TimingMode: chModel.TimingModeDeadline, DueAt: &due, PerformedAt: due.Add(9 * time.Hour), Outcome: ScoreOutcomeRecovered},
			expected: chModel.ScoreResult{BasePoints: 10, TimingAdjustment: -7, RecoveryPoints: 3, Total: 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, ScoreTaskResult(test.input))
		})
	}
}

func TestCompletionOutcomeOnlyRecoversMatchingMiss(t *testing.T) {
	due := time.Date(2026, 8, 26, 18, 0, 0, 0, time.UTC)
	otherDue := due.Add(24 * time.Hour)
	require.Equal(t, ScoreOutcomeCompleted, completionOutcome(nil, &due))
	require.Equal(t, ScoreOutcomeCompleted, completionOutcome([]*chModel.ChoreHistory{{Status: chModel.ChoreHistoryStatusMissed, DueDate: &otherDue}}, &due))
	require.Equal(t, ScoreOutcomeRecovered, completionOutcome([]*chModel.ChoreHistory{{Status: chModel.ChoreHistoryStatusMissed, DueDate: &due}}, &due))
}

func TestMissBoundaryUsesTimingMode(t *testing.T) {
	due := time.Date(2026, 8, 26, 18, 0, 0, 0, time.UTC)
	points := 10
	deadline := &chModel.Chore{Points: &points, TimingMode: chModel.TimingModeDeadline, NextDueDate: &due}
	require.Equal(t, due.Add(8*time.Hour), *missBoundary(deadline))

	today := &chModel.Chore{Points: &points, TimingMode: chModel.TimingModeToday, NextDueDate: &due, FrequencyMetadataV2: &chModel.FrequencyMetadata{Timezone: "America/New_York"}}
	require.Equal(t, time.Date(2026, 8, 27, 3, 59, 59, 999999999, time.UTC), *missBoundary(today))

	untimed := &chModel.Chore{Points: &points, TimingMode: chModel.TimingModeUntimed, NextDueDate: &due}
	require.Nil(t, missBoundary(untimed))
}
