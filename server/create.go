package server

import (
	pb "github.com/baopham/goproto/reminder"
	"github.com/baopham/reminder/db"
	"github.com/lileio/lile/pubsub"
	context "golang.org/x/net/context"
)

func (s ReminderServer) Create(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	session, err := db.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	id, err := db.CreateReminder(r, session)
	response := &pb.CreateResponse{Id: id}
	go pubsub.Publish(ctx, "summarizer.start", r)
	return response, err
}
