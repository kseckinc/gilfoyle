package cmd

import (
	"fmt"
	"github.com/dreamvo/gilfoyle/config"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "gilfoyle [COMMANDS] [OPTIONS]",
	Short:   "Cloud-native media streaming server",
	Long:    "Gilfoyle is a web application from the Dreamvo project that runs a self-hosted media streaming server.",
	Example: "gilfoyle serve -p 8080 --config ./gilfoyle.yml",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")
}

func initConfig() {
	if cfgFile != "" {
		err := config.New(cfgFile)
		if err != nil {
			panic(err)
		}
		return
	}

	err := config.New()
	if err != nil {
		panic(err)
	}
}

// Execute is a function that executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
