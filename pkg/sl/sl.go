package sl

import (
	"fmt"
	"log/slog"
)

func Err(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}

func Attr(err error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.StringValue(err.Error())}
}
