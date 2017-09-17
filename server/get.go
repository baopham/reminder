package server

import (
	pb "github.com/baopham/goproto/reminder"
	"github.com/baopham/reminder/db"
	context "golang.org/x/net/context"
)

func (s ReminderServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	session, err := db.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	reminder, err := db.FindReminder(r.Id, session)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Reminder: reminder}, nil
}
