/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"leoagomes/todo/internal/cli/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A tool for managing project-wide tasks",
	Long: `todo is a tool for managing project-wide tasks. It is a command line tool 
that allows you to create, update, and delete tasks. It is a tool that allows 
you to manage your tasks in a terminal.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags are global for the application
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .todo.yaml in current directory or $HOME)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var err error
	cfg, err = config.Load(cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
}

// GetConfig returns the loaded configuration
func GetConfig() *config.Config {
	return cfg
}
