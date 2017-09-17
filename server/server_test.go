package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/baopham/reminder"
	"github.com/lileio/lile"
)

var s = ReminderServer{}
var cli reminder.ReminderClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		reminder.RegisterReminderServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = reminder.NewReminderClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
