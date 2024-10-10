package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  *slog.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crib",
	Short: "A brief description of your application",
	Long: `CRIB stands for "Chainlink Running-in-a-Box".
	
CRIB is tooling that enables CLL developers to quickly spin up ephemeral development 
and/or testing environments that closely mimic a product’s staging environment with 
all the required Chainlink dependencies.`,

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

//nolint:gochecknoinits
func init() {
	cobra.OnInitialize(ensureRunningInAProductDir, initConfig, initLogger)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".env", "config file")
	rootCmd.PersistentFlags().StringP("log-level", "", "info", "Log level (debug, info, warn, error)")

	// Bind the viper flag to the cobra flag (we can safely ignore the error here)
	_ = viper.BindPFlag("log_level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.SetDefault("log_level", "info")
}

func ensureRunningInAProductDir() {
	repoRoot, err := utils.GetGitTopLevelDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to determine git repository root path: %v\n", err)
		os.Exit(1)
	}

	productsDir := filepath.Join(repoRoot, "deployments")
	availableProducts, err := utils.ListFiles(productsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list products dir %s: %v\n", productsDir, err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	relPath, err := filepath.Rel(dir, repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get relative path: %v\n", err)
		os.Exit(1)
	}

	if !slices.Contains(availableProducts, filepath.Base(dir)) {
		fmt.Fprintf(os.Stderr, "Current working directory is not a product directory. Make sure you cd into one of %s/deployments/%v and try again\n", relPath, availableProducts)
		os.Exit(1)
	}
}

// initConfig reads in a config file or initializes a new one from an example
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// TODO: evaluate using a config file in the user's home directory
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")
	}

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
		return
	}

	fmt.Fprintln(os.Stderr, "ERROR: ", err)

	// offer the user to use the predefined .env file from .env.example
	validChoices := []string{"yes", "y", "no", "n"}
	userChoice := utils.PresentPrompt("CRIB deployment requires several environment variables. Since you don’t have a custom '.env' file set up, would you like to use the predefined '.env' file instead? (yes/no): ", validChoices)
	if userChoice == "yes" || userChoice == "y" {
		err := utils.CopyFile(".env.example", ".env")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to copy .env.example to .env: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Exiting without providing any '.env' file.")
		os.Exit(1)
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read in the config file: ", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}

func initLogger() {
	requestedLogLevel := viper.GetString("log_level")
	var slogLevel slog.Level
	var w io.Writer
	switch requestedLogLevel {
	case "debug":
		slogLevel = slog.LevelDebug
		w = os.Stdout
	case "info":
		slogLevel = slog.LevelInfo
		w = os.Stdout
	case "warn":
		slogLevel = slog.LevelWarn
		w = os.Stderr
	case "error":
		slogLevel = slog.LevelError
		w = os.Stderr
	default:
		fmt.Fprintf(os.Stderr, "Invalid log level %s", requestedLogLevel)
		os.Exit(1)
	}
	logger = slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: slogLevel}))
	logger.Debug("Debug mode enabled")
}
