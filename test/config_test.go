package test

import (
	"github.com/goccy/go-json"
	"golang.org/x/exp/slog"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))

}

type Status int

const (
	Running Status = 1
	Except  Status = 2
	OffLine Status = 3
)

func (s Status) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "OffLine"
	}
	return json.Marshal(str)
}
