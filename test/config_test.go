package test

import (
	"golang.org/x/exp/slog"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))

}
