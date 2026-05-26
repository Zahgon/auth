package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/supabase/auth/internal/conf"
)

var autoconfirm, isAdmin bool
var audience string

func getAudience(c *conf.GlobalConfiguration) string { _ = "STUB: not implemented"; return "" }

func adminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

var adminCreateUserCmd = cobra.Command{
	Use: "createuser",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logrus.Fatal("Not enough arguments to createuser command. Expected at least email and password values")
			return
		}

		execWithConfigAndArgs(cmd, adminCreateUser, args)
	},
}

var adminDeleteUserCmd = cobra.Command{
	Use: "deleteuser",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatal("Not enough arguments to deleteuser command. Expected at least ID or email")
			return
		}

		execWithConfigAndArgs(cmd, adminDeleteUser, args)
	},
}

func adminCreateUser(config *conf.GlobalConfiguration, args []string) {
	_ = "STUB: not implemented"
	return
}

func adminDeleteUser(config *conf.GlobalConfiguration, args []string) {
	_ = "STUB: not implemented"
	return
}
