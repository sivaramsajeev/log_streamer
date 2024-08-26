package cmd

import (
	"github.com/sivaramsajeev/log_streamer/consumer"
	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive messages from streamer topic",
	Long:  `Receive messages to kafka after ensuring proper configs are in place`,
	Run: func(cmd *cobra.Command, args []string) {
		consumer.NewKafkaReceiver().Receive()
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}
