package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Timofey-Grishechko/golang-todoapp/internal/core/domain"
	core_logger "github.com/Timofey-Grishechko/golang-todoapp/internal/core/logger"
	core_http_request "github.com/Timofey-Grishechko/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/Timofey-Grishechko/golang-todoapp/internal/core/transport/http/response"
)

type GetStatisticsRespnse struct {
	TaskCreated               int      `json:"task_created"`
	TaskCompleted             int      `json:"task_completed"`
	TaskCompletedRate         *float64 `json:"task_completed_rate"`
	TaskAverageCompletionTime *string  `json:"task_average_completion_time"`
}

func (h *StatisticsHttpHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIDFromToQuertyParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/from/to query params",
		)

		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)

		return
	}

	response := toDTOFromDomain(statistics)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsRespnse {
	var avgTime *string
	if statistics.TaskAverageCompletionTime != nil {
		duration := statistics.TaskAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsRespnse{
		TaskCreated:               statistics.TaskCreated,
		TaskCompleted:             statistics.TaskCompleted,
		TaskCompletedRate:         statistics.TaskCompletedRate,
		TaskAverageCompletionTime: avgTime,
	}
}

func getUserIDFromToQuertyParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParanKey   = "from"
		toQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParanKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}
