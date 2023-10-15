package main

import (
	"github.com/spf13/cobra"
	"os"
	"wasm-forge/internal/app/api"
)

var rootCmd = &cobra.Command{
	Use:   "wasm-forge",
	Short: "Build wasm for any language",
	Run: func(cmd *cobra.Command, args []string) {
		if runMode == "bin" {
			// TODO
		} else if runMode == "api" {
			api.Start()
		}
	},
}

var runMode string

func init() {
	rootCmd.Flags().StringVarP(&runMode, "run-mode", "m", "bin", "run mode [bin, api]")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
