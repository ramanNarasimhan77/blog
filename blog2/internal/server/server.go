package server

import (
	"context"
	"time"

	"github.com/arkadyb/demos/blog2/proto/reminder/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyServer struct {
}

func (s *MyServer) ScheduleReminder(ctx context.Context, req *reminder.ScheduleReminderRequest) (*reminder.ScheduleReminderResponse, error) {
	if req.When == nil {
		return nil, status.Error(codes.InvalidArgument, "when cant be nil")
	}

	when, err := ptypes.Timestamp(req.GetWhen())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "when is invalid")
	}

	if when.Before(time.Now()) {
		return nil, status.Error(codes.InvalidArgument, "when should be in the future")
	}

	go func(dur time.Duration) {
		timer := time.NewTimer(dur)
		<-timer.C
		log.Info("Timer time!")
	}(when.Sub(time.Now()))

	return &reminder.ScheduleReminderResponse{
		Id: uuid.New().String(),
	}, nil
}
