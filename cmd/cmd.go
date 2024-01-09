package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "kafkid",
}

func init() {
	rootCmd.AddCommand(RunPubisher())
	rootCmd.AddCommand(RunConsumer())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
