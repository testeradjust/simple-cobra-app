/*
Copyright © 2020 Serj Tubin <stubin87@google.com.com>

*/
package cmd

import (
	"fmt"

	"github.com/testeradjust/simple-cobra-app/internal"

	"github.com/spf13/cobra"
)

// aeroCmd represents the aero command
var aeroCmd = &cobra.Command{
	Use:   "aero",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating aero ...")
		dummyAero, err := internal.NewDummyAero("172.28.128.3", 3000)
		if err != nil {
			fmt.Printf("cannot create dummy aero: %s\n", err)
			return
		}

		fmt.Println("aero created, testing connection ...")
		if err := dummyAero.TestConnection(); err != nil {
			fmt.Printf("aero test failed: %s\n", err)
			return
		}

		fmt.Println("aero done")
	},
}

func init() {
	rootCmd.AddCommand(aeroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aeroCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aeroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
