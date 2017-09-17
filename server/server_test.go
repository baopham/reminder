package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/baopham/goproto/reminder"
	"github.com/lileio/lile"
)

var s = ReminderServer{}
var cli pb.ReminderClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		pb.RegisterReminderServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = pb.NewReminderClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
