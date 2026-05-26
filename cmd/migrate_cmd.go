package cmd

import (
	"embed"

	"github.com/spf13/cobra"
)

var EmbeddedMigrations embed.FS

var migrateCmd = cobra.Command{
	Use:  "migrate",
	Long: "Migrate database strucutures. This will create new tables and add missing columns and indexes.",
	Run:  migrate,
}

func migrate(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// Set to true to display query info

// Hide pop migration logging

// turn off schema dump
