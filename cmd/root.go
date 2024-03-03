/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "a",
		Short: "a: an unofficial ansible experiment",
	}
	logLevel string
	logger   *log.Logger
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "info", "Log level (debug, info, warn, error, fatal, panic)")

	cobra.OnInitialize(configureLogger)
}

func configureLogger() {
	programLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("invalid log level", "error", err)
	}

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
	logger.SetLevel(programLevel)
	logger.Info("Welcome to a - an unofficial ansible experiment!")
}
