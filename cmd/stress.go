/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/leonardfreitas/stress_test_go/internal/commands"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
	timeout     int
)

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Run stress test over a given URL",
	Long:  "Run stress test over a given URL. It will make a number of requests to the given URL and measure the status codes",
	Run: func(cmd *cobra.Command, args []string) {
		command := commands.NewStressCommand(concurrency, requests, url, timeout)
		command.Run()
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)

	stressCmd.Flags().StringVar(&url, "url", "", "URL to be tested")
	stressCmd.MarkFlagRequired("url")
	stressCmd.Flags().IntVar(&requests, "requests", 1000, "Number of requests")
	stressCmd.Flags().IntVar(&concurrency, "concurrency", 10, "Number of concurrent requests")
	stressCmd.Flags().IntVar(&timeout, "timeout", 3, "Timeout in seconds")
}
