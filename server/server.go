package server

import (
	pb "github.com/baopham/goproto/reminder"
)

type ReminderServer struct {
	pb.ReminderServiceServer
}
