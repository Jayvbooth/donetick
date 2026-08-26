package chore

import (
	"context"
	"errors"
	"sync"
	"time"

	chRepo "donetick.com/core/internal/chore/repo"
	"donetick.com/core/logging"
)

const missedScoringInterval = 15 * time.Minute

type MissedScoringService struct {
	repo     *chRepo.ChoreRepository
	stop     chan struct{}
	stopOnce sync.Once
}

func NewMissedScoringService(repo *chRepo.ChoreRepository) *MissedScoringService {
	return &MissedScoringService{repo: repo, stop: make(chan struct{})}
}

func (s *MissedScoringService) Start(c context.Context) {
	go func() {
		s.record(c, time.Now().UTC())
		ticker := time.NewTicker(missedScoringInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				s.record(c, time.Now().UTC())
			case <-s.stop:
				return
			}
		}
	}()
}

func (s *MissedScoringService) Stop() {
	s.stopOnce.Do(func() { close(s.stop) })
}

func (s *MissedScoringService) record(c context.Context, now time.Time) {
	chores, err := s.repo.GetMissedScoreCandidates(c, now)
	if err != nil {
		logging.FromContext(c).Error("Failed to load missed scoring candidates", "error", err)
		return
	}
	for _, chore := range chores {
		boundary := missBoundary(chore)
		if boundary == nil || now.Before(*boundary) {
			continue
		}
		score := scoreForChore(chore, now, ScoreOutcomeMissed)
		if err := s.repo.RecordMissedScore(c, chore, now, score); err != nil && !errors.Is(err, chRepo.ErrMissAlreadyScored) && !errors.Is(err, chRepo.ErrChoreChanged) {
			logging.FromContext(c).Error("Failed to record missed score", "error", err, "choreID", chore.ID)
		}
	}
}
