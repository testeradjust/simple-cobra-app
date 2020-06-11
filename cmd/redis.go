/*
Copyright © 2020 Serj Tubin <stubin87@google.com.com>

*/
package cmd

import (
	"fmt"

	"github.com/testeradjust/simple-cobra-app/internal"

	"github.com/spf13/cobra"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("redis called ...")
		dummyRedis, err := internal.NewDummyRedis("localhost", 6379)
		if err != nil {
			fmt.Printf("too bad, failed to create redis: %s\n", err)
			return
		}

		fmt.Println("redis created, testing connection ...")
		if err := dummyRedis.TestConnection(); err != nil {
			fmt.Printf("hmmm, failed to test redis connection: %s\n", err)
			return
		}

		fmt.Println("redis, done")
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// redisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// redisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
