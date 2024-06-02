package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gov",
	Short: "Run test and make coverage report in go multi module project",
	Long:  "Run test and make coverage report in go multi module project",
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running command...")
		nameVar, _ := cmd.Flags().GetString("name")
		name = nameVar
		fmt.Println(printHelloWorld())
	},
}
var name string

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize()
	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVar(&name, "name", "Me", "name to print in the message")
}

func printHelloWorld() string {
	return fmt.Sprintf("Hello, %s", name)
}
