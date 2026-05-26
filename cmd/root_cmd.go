package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/supabase/auth/internal/conf"
)

var (
	configFile = ""
	watchDir   = ""
)

var rootCmd = cobra.Command{
	Use: "gotrue",
	Run: func(cmd *cobra.Command, args []string) {
		migrate(cmd, args)
		serve(cmd.Context())
	},
}

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func loadGlobalConfig(ctx context.Context) *conf.GlobalConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func execWithConfigAndArgs(cmd *cobra.Command, fn func(config *conf.GlobalConfiguration, args []string), args []string) {
	_ = "STUB: not implemented"
	return
}
