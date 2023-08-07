package cli

import (
	"github.com/spf13/cobra"
	"os"
)

var wad_export_path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wad2linkedin",
	Short: "Small cli tool to help connecting with people who we met at WAD",
	Long: `Right now it only supports the show command, which opens the browser with the search results lead by lead

			Example:
			wad2linkedin show -file ./contacts.xlsx`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.arango.yaml)")
	rootCmd.PersistentFlags().StringVar(&wad_export_path, "file", "", "The path to the WAD export file.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
