package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/smartcontractkit/devenv/ccipv17/services"
	"github.com/spf13/cobra"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"

	ccipv17 "github.com/smartcontractkit/devenv/ccipv17"
)

const (
	LocalCCIPv2Dashboard = "http://localhost:3000/d/f8a04cef-653f-46d3-86df-87c532300672/datafeedsv1-soak-test?orgId=1&refresh=5s"
)

var rootCmd = &cobra.Command{
	Use:   "ccip",
	Short: "A CCIP local environment tool",
}

var reconfigureCmd = &cobra.Command{
	Use:     "reconfigure",
	Aliases: []string{"r"},
	Short:   "Reconfigure development environment, remove apps and apply new configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, _ := rootCmd.Flags().GetString("config")
		framework.L.Info().Str("Config", configFile).Msg("Reconfiguring development environment")
		_ = os.Setenv("CTF_CONFIGS", configFile)
		_ = os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
		framework.L.Info().Msg("Tearing down the development environment")
		err := framework.RemoveTestContainers()
		if err != nil {
			return fmt.Errorf("failed to clean Docker resources: %w", err)
		}
		_, err = ccipv17.NewEnvironment(true, "")
		return err
	},
}

var upCmd = &cobra.Command{
	Use:     "up",
	Aliases: []string{"u"},
	Short:   "Spin up the development environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, _ := rootCmd.Flags().GetString("config")
		framework.L.Info().Str("Config", configFile).Msg("Creating development environment")
		_ = os.Setenv("CTF_CONFIGS", configFile)
		_ = os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
		_, err := ccipv17.NewEnvironment(true, "")
		if err != nil {
			return err
		}
		return nil
	},
}

var downCmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"d"},
	Short:   "Tear down the development environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		// namespace, _ := rootCmd.Flags().GetString("namespace")
		if !remote {
			framework.L.Info().Msg("Tearing down the development environment")
			err := framework.RemoveTestContainers()
			if err != nil {
				return fmt.Errorf("failed to clean Docker resources: %w", err)
			}
		}
		// TODO: add CRIB SDK integration
		return nil
	},
}

var bsCmd = &cobra.Command{
	Use:   "bs",
	Short: "Manage the Blockscout EVM block explorer",
	Long:  "Spin up or down the Blockscout EVM block explorer",
}

var bsUpCmd = &cobra.Command{
	Use:     "up",
	Aliases: []string{"u"},
	Short:   "Spin up Blockscout EVM block explorer",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		url, _ := bsCmd.Flags().GetString("url")
		if remote {
			return fmt.Errorf("remote mode: %v, Blockscout can only be used in 'local' mode", remote)
		}
		return framework.BlockScoutUp(url)
	},
}

var bsDownCmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"d"},
	Short:   "Spin down Blockscout EVM block explorer",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		url, _ := bsCmd.Flags().GetString("url")
		if remote {
			return fmt.Errorf("remote mode: %v, Blockscout can only be used in 'local' mode", remote)
		}
		return framework.BlockScoutDown(url)
	},
}

var bsRestartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"r"},
	Short:   "Restart the Blockscout EVM block explorer",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		url, _ := bsCmd.Flags().GetString("url")
		if !remote {
			return fmt.Errorf("remote mode: %v, Blockscout can only be used in 'local' mode", remote)
		}
		if err := framework.BlockScoutDown(url); err != nil {
			return err
		}
		return framework.BlockScoutUp(url)
	},
}

var obsCmd = &cobra.Command{
	Use:   "obs",
	Short: "Manage the observability stack",
	Long:  "Spin up or down the observability stack with subcommands 'up' and 'down'",
}

var obsUpCmd = &cobra.Command{
	Use:     "up",
	Aliases: []string{"u"},
	Short:   "Spin up the observability stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		if remote {
			return fmt.Errorf("remote mode: %v, local observability stack can only be used in 'local' mode", remote)
		}
		if err := framework.ObservabilityUpFull(); err != nil {
			return fmt.Errorf("observability up failed: %w", err)
		}
		ccipv17.Plog.Info().Msgf("CCIPv17 Dashboard: %s", LocalCCIPv2Dashboard)
		return nil
	},
}

var obsDownCmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"d"},
	Short:   "Spin down the observability stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := rootCmd.Flags().GetBool("remote")
		if remote {
			return fmt.Errorf("remote mode: %v, local observability stack can only be used in 'local' mode", remote)
		}
		return framework.ObservabilityDown()
	},
}

var obsRestartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"r"},
	Short:   "Restart the observability stack (data wipe)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := framework.ObservabilityDown(); err != nil {
			return fmt.Errorf("observability down failed: %w", err)
		}
		if err := framework.ObservabilityUpFull(); err != nil {
			return fmt.Errorf("observability up failed: %w", err)
		}
		ccipv17.Plog.Info().Msgf("CCIPv17 Dashboard: %s", LocalCCIPv2Dashboard)
		return nil
	},
}

var indexerDBShell = &cobra.Command{
	Use:     "indexer-db",
	Aliases: []string{"idb"},
	Short:   "Inspect Indexer Database",
	RunE: func(cmd *cobra.Command, args []string) error {
		psqlPath, err := exec.LookPath("psql")
		if err != nil {
			return fmt.Errorf("psql not found in PATH, are you inside 'nix develop' shell?: %w", err)
		}
		psqlArgs := []string{
			"psql",
			services.DefaultDBConnectionString,
		}
		if len(args) > 0 {
			psqlArgs = append(psqlArgs, args...)
		}
		env := syscall.Environ()
		return syscall.Exec(psqlPath, psqlArgs, env)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "env.toml,overrides.toml", "Configuration file for the environment")
	rootCmd.PersistentFlags().StringP("blockscout_url", "u", "http://host.docker.internal:8545", "EVM RPC node URL")

	bsCmd.AddCommand(bsUpCmd)
	bsCmd.AddCommand(bsDownCmd)
	bsCmd.AddCommand(bsRestartCmd)
	rootCmd.AddCommand(bsCmd)

	obsCmd.AddCommand(obsRestartCmd)
	obsCmd.AddCommand(obsUpCmd)
	obsCmd.AddCommand(obsDownCmd)
	rootCmd.AddCommand(obsCmd)

	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(reconfigureCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(indexerDBShell)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		ccipv17.Plog.Err(err).Send()
		os.Exit(1)
	}
}
