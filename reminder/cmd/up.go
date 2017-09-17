package cmd

import (
	"github.com/baopham/reminder/subscribers"
	"github.com/lileio/lile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC and pubub subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logrus.Fatal(lile.Serve())

		}()

		lile.Subscribe(&subscribers.ReminderServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
