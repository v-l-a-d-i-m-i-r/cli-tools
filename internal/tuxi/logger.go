package tuxi

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	logLevelEnv = "TUXI_LOG_LEVEL"
	logFileEnv  = "TUXI_LOG_FILE"
)

// newFileLogger builds a slog.Logger configured from the TUXI_LOG_LEVEL and
// TUXI_LOG_FILE environment variables. TUXI_LOG_LEVEL accepts debug, info,
// warn, error (case-insensitive). TUXI_LOG_FILE, if set, is opened for
// append and used as the log output instead of stderr. If neither variable
// is set, the returned logger discards everything, so callers can use it
// unconditionally.
func newFileLogger() (*slog.Logger, error) {
	_, levelSet := os.LookupEnv(logLevelEnv)
	_, fileSet := os.LookupEnv(logFileEnv)

	if !levelSet && !fileSet {
		return slog.New(slog.DiscardHandler), nil
	}

	level := parseLevel(os.Getenv(logLevelEnv))

	out, err := resolveOutput(os.Getenv(logFileEnv))
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(out, &slog.HandlerOptions{Level: level})

	return slog.New(handler), nil
}

func parseLevel(raw string) slog.Level {
	switch strings.ToLower(raw) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func resolveOutput(path string) (io.Writer, error) {
	if path == "" {
		return os.Stderr, nil
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
