package cmd

import (
	"github.com/sivaramsajeev/log_streamer/producer"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a Kafka topic",
	Long:  `Send messages to kafka after ensuring proper configs are in place`,
	Run: func(cmd *cobra.Command, args []string) {
		producer.NewKafkaSender().Send()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
