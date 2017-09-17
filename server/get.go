package server

import (
	"github.com/baopham/reminder"
	"github.com/baopham/reminder/db"
	context "golang.org/x/net/context"
)

func (s ReminderServer) Get(ctx context.Context, r *reminder.GetRequest) (*reminder.ReminderResponse, error) {
	session, err := db.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return db.FindReminder(r.Id, session)
}
