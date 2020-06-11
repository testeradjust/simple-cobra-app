/*
Copyright © 2020 Serj Tubin <stubin87@google.com.com>

*/
package cmd

import (
	"fmt"

	"github.com/testeradjust/simple-cobra-app/internal"

	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kafka called ...")

		brokers := []string{"localhost:9092"}
		dummyKafka, err := internal.NewDummyKafka(brokers)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("sending test message ...")
		dummyKafka.ProduceTestMessage("bla bla message :8")

		fmt.Println("done, closing ...")
		if err := dummyKafka.Close(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("kafka all done")
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
