package server

import (
	"github.com/baopham/reminder"
	"github.com/baopham/reminder/db"
	"github.com/lileio/lile/pubsub"
	context "golang.org/x/net/context"
)

func (s ReminderServer) Create(ctx context.Context, r *reminder.CreateRequest) (*reminder.CreateResponse, error) {
	session, err := db.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	id, err := db.CreateReminder(r, session)
	response := &reminder.CreateResponse{Id: id}
	pubsub.Publish(ctx, "reminder.created", response)
	return response, err
}
