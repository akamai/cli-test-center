package cmd

import (
	"errors"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
	Use: externalconstant.RootCommandUse,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if forceColor {
			color.NoColor = false
		}

		var err error
		config, err = edgegrid.InitEdgeRc(edgeRcPath, edgeRcSection)
		if err != nil {
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.InitEdgeRc), internalconstant.ExitStatusCode1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(Version string) {
	rootCmd.Version = Version
	code := internalconstant.ExitStatusCode0
	if err := rootCmd.Execute(); err != nil {
		switch err.Error() {
		case "2":
			// For rootCmd and Subcommand invalid flags
			code = internalconstant.ExitStatusCode2
			break
		default:
			util.PrintError(err.Error() + "\n")
			rootCmd.Println("\n" + rootCmd.UsageString())
			// For rootCmd wrong arguments
			code = internalconstant.ExitStatusCode2
			break
		}
	}
	os.Exit(code)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Remove this if we choose to offer a completion command
	isStandardInputAvailable, jsonData = util.ReadStdin(rootCmd)

	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	// AKATEST-9393: Print flag error in red color and usage without color
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	// Setting the default values for the global flags from environment variable.
	defaultEdgercPath := os.Getenv(internalconstant.DefaultEdgercPathKey)
	defaultEdgercSection := os.Getenv(internalconstant.DefaultEdgercSectionKey)
	globalJsonFlag, _ := strconv.ParseBool(os.Getenv(internalconstant.DefaultJsonOutputKey))

	if defaultEdgercPath == internalconstant.Empty {
		if home, err := homedir.Dir(); err == nil {
			defaultEdgercPath = filepath.Join(home, internalconstant.EdgercFileNameDefaultValue)
		}
	}

	if defaultEdgercSection == internalconstant.Empty {
		defaultEdgercSection = internalconstant.SectionDefaultValue
	}

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		util.PrintError(err.Error() + "\n\n")
		cmd.Println(cmd.UsageString())
		return errors.New(strconv.Itoa(internalconstant.ExitStatusCode2))
	})

	rootCmd.Short = util.GetMessageForKey(rootCmd, internalconstant.Short)
	rootCmd.Long = util.GetMessageForKey(rootCmd, internalconstant.Long)

	rootCmd.PersistentFlags().StringVarP(&edgeRcPath, externalconstant.FlagEdgerc, externalconstant.FlagEdgercShortHand, defaultEdgercPath, util.GetMessageForKey(rootCmd, externalconstant.FlagEdgerc))
	rootCmd.PersistentFlags().StringVarP(&edgeRcSection, externalconstant.FlagSection, externalconstant.FlagSectionShortHand, defaultEdgercSection, util.GetMessageForKey(rootCmd, externalconstant.FlagSection))
	rootCmd.PersistentFlags().StringVar(&accountSwitchKey, externalconstant.FlagAccountKey, internalconstant.Empty, util.GetMessageForKey(rootCmd, externalconstant.FlagAccountKey))
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, externalconstant.FlagJson, globalJsonFlag, util.GetMessageForKey(rootCmd, externalconstant.FlagJson))
	rootCmd.PersistentFlags().BoolVar(&forceColor, externalconstant.FlagForceColor, false, util.GetMessageForKey(rootCmd, externalconstant.FlagForceColor))
	rootCmd.Flags().BoolP(externalconstant.FlagHelp, externalconstant.FlagHelpShortHand, false, util.GetMessageForKey(rootCmd, externalconstant.FlagHelp))
	rootCmd.Flags().Bool(externalconstant.FlagVersion, false, util.GetMessageForKey(rootCmd, externalconstant.FlagVersion))
}
