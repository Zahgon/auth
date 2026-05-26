package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start API server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func serve(ctx context.Context) { _ = "STUB: not implemented"; return }

// Include serve ctx which carries cancelation signals so DialContext does
// not hang indefinitely at startup.

// Add the base context to the db, this is so during the shutdown sequence
// the DB will be available while connections drain.

// Do not return to caller until this goroutine is done.

// to mitigate a Slowloris attack

// Work exits when ctx is done as in-flight requests do not depend
// on it. If they do in the future this should be baseCtx instead.

// Update the previous limiter with the latest config

// Create a new API version with the updated config.

// Create a new mailer with existing template cache.

// Persist existing rate limiters.

// Assign this config as the latest configuration

// When config is updated we notify the apiworker.

// Update previous limiter

// #nosec G118 -- Cleanup goroutine intentionally outlives the request; context.Background() is required for shutdown after parent context is cancelled.

// This must be done after httpSrv exits, otherwise you may potentially
// have 1 or more inflight http requests blocked until the shutdownCtx
// is canceled.

// #nosec G115
