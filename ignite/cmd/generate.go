package ignitecmd

import (
	"github.com/spf13/cobra"

	"github.com/ignite/cli/ignite/pkg/cliui"
)

// NewGenerate returns a command that groups code generation related sub commands.
func NewGenerate() *cobra.Command {
	c := &cobra.Command{
		Use:   "generate [command]",
		Short: "Generate clients, API docs from source code",
		Long: `Generate clients, API docs from source code.

Such as compiling protocol buffer files into Go or implement particular
functionality, for example, generating an OpenAPI spec.

Produced source code can be regenerated by running a command again and is not
meant to be edited by hand.
`,
		Aliases:           []string{"g"},
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: generatePreRunHandler,
	}

	flagSetPath(c)
	flagSetClearCache(c)
	c.AddCommand(NewGenerateGo())
	c.AddCommand(NewGenerateTSClient())
	c.AddCommand(NewGenerateVuex())
	c.AddCommand(NewGenerateComposables())
	c.AddCommand(NewGenerateHooks())
	c.AddCommand(NewGenerateOpenAPI())

	return c
}

func generatePreRunHandler(cmd *cobra.Command, _ []string) error {
	session := cliui.New()
	defer session.End()

	return toolsMigrationPreRunHandler(cmd, session)
}
