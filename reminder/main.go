package main

import (
	pb "github.com/baopham/goproto/reminder"
	"github.com/baopham/reminder/reminder/cmd"
	"github.com/baopham/reminder/server"
	"github.com/lileio/lile"
	"github.com/lileio/lile/fromenv"
	"github.com/lileio/lile/pubsub"
	"google.golang.org/grpc"
)

func main() {
	s := &server.ReminderServer{}

	lile.Name("reminder")
	lile.Server(func(g *grpc.Server) {
		pb.RegisterReminderServiceServer(g, s)
	})

	service := lile.GlobalService()
	pubsub.SetClient(&pubsub.Client{
		Provider: fromenv.PubSubProvider(service.Name),
	})

	cmd.Execute()
}
