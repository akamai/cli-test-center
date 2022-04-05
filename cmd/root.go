package cmd

import (
	"errors"
	"os"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/akamai/cli-test-center/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	edgeRcPath               string
	edgeRcSection            string
	accountSwitchKey         string
	jsonOutput               bool
	forceColor               bool
	config                   edgegrid.Config
	jsonData                 []byte
	isStandardInputAvailable bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: RootCommandUse,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if forceColor {
			color.NoColor = false
		}

		var err error
		config, err = edgegrid.InitEdgeRc(edgeRcPath, edgeRcSection)
		if err != nil {
			internal.AbortWithExitCode(internal.GetGlobalErrorMessage("initEdgeRc"), internal.ExitStatusCode1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(Version string) {
	rootCmd.Version = Version
	code := internal.ExitStatusCode0
	if err := rootCmd.Execute(); err != nil {
		switch err.Error() {
		case "2":
			// For rootCmd and Subcommand invalid flags
			code = internal.ExitStatusCode2
			break
		default:
			internal.PrintError(err.Error() + "\n")
			rootCmd.Println("\n" + rootCmd.UsageString())
			// For rootCmd wrong arguments
			code = internal.ExitStatusCode2
			break
		}
	}
	os.Exit(code)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Remove this if we choose to offer a completion command
	isStandardInputAvailable, jsonData = internal.ReadStdin(rootCmd)

	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	// AKATEST-9393: Print flag error in red color and usage without color
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		internal.PrintError(err.Error() + "\n\n")
		cmd.Println(cmd.UsageString())
		return errors.New(strconv.Itoa(internal.ExitStatusCode2))
	})

	rootCmd.Short = internal.GetMessageForKey(rootCmd, internal.Short)
	rootCmd.Long = internal.GetMessageForKey(rootCmd, internal.Long)

	rootCmd.PersistentFlags().StringVar(&edgeRcPath, FlagEdgerc, FlagEdgercDefaultValue, internal.GetMessageForKey(rootCmd, FlagEdgerc))
	rootCmd.PersistentFlags().StringVar(&edgeRcSection, FlagSection, FlagSectionDefaultValue, internal.GetMessageForKey(rootCmd, FlagSection))
	rootCmd.PersistentFlags().StringVar(&accountSwitchKey, FlagAccountKey, "", internal.GetMessageForKey(rootCmd, FlagAccountKey))
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, FlagJson, false, internal.GetMessageForKey(rootCmd, FlagJson))
	rootCmd.PersistentFlags().BoolVar(&forceColor, FlagForceColor, false, internal.GetMessageForKey(rootCmd, FlagForceColor))
	rootCmd.Flags().BoolP(FlagHelp, FlagHelpShortHand, false, internal.GetMessageForKey(rootCmd, FlagHelp))
	rootCmd.Flags().BoolP(FlagVersion, FlagVersionShortHand, false, internal.GetMessageForKey(rootCmd, FlagVersion))
}
