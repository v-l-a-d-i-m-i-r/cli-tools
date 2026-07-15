// Package tuxi implements the tuxi multi-command CLI.
package tuxi

import (
	"context"
	"log/slog"

	"github.com/spf13/cobra"
)

type loggerContextKey struct{}

// NewRootCmd constructs the root tuxi command.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "tuxi",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			log, err := newFileLogger()
			if err != nil {
				return err
			}

			cmd.SetContext(context.WithValue(cmd.Context(), loggerContextKey{}, log))

			return nil
		},
	}

	rootCmd.AddCommand(newWindowStatusFormatCmd())

	return rootCmd
}

func loggerFromContext(ctx context.Context) *slog.Logger {
	if log, ok := ctx.Value(loggerContextKey{}).(*slog.Logger); ok {
		return log
	}

	return slog.New(slog.DiscardHandler)
}
