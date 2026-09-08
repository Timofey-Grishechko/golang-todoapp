package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Timofey-Grishechko/golang-todoapp/internal/core/domain"
	core_errors "github.com/Timofey-Grishechko/golang-todoapp/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get tasks from repository: %w", err)
	}

	statistics := culcStatistics(tasks)

	return statistics, nil
}

func culcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.NewStatistics(0, 0, nil, nil)

	}

	taskCreated := len(tasks)

	taskCompleted := 0
	var totalCompletionDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			taskCompleted++
		}

		complitionDuration := task.CompletionDuration()
		if complitionDuration != nil {
			totalCompletionDuration += *complitionDuration
		}
	}

	taskCompletedRate := float64(taskCompleted) / float64(taskCreated) * 100

	var taskAverageCompletionTime *time.Duration
	if taskCompleted > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(taskCompleted)

		taskAverageCompletionTime = &avg
	}

	return domain.NewStatistics(
		taskCreated,
		taskCompleted,
		&taskCompletedRate,
		taskAverageCompletionTime,
	)
}
