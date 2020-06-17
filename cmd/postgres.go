/*
Copyright © 2020 Serj Tubin <stubin87@google.com.com>

*/
package cmd

import (
	"fmt"

	"github.com/testeradjust/simple-cobra-app/internal"

	"github.com/spf13/cobra"
)

var dbUser string
var dbName string
var dbPassword string

// postgresCmd represents the postgres command
var postgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("postgres called ...")
		fmt.Printf("user: %s, psw: %s, db: %s\n", dbUser, dbPassword, dbName)
		dummyPostgres, err := internal.NewDummyPostgres("localhost", dbUser, dbPassword, dbName, 5432)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("postgres conn created, testing connection ...")
		if err := dummyPostgres.TestConnection(); err != nil {
			fmt.Printf("failed to ping ps db: %s\n", err)
			return
		}

		fmt.Println("all OK")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbUser, "user", "u", "dummydb", "name of db user")
	rootCmd.PersistentFlags().StringVarP(&dbName, "name", "n", "dummydb", "name of db")
	rootCmd.PersistentFlags().StringVarP(&dbPassword, "password", "p", "dummydbps", "user db password")

	rootCmd.AddCommand(postgresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postgresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postgresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
