package main

import (
	"fmt"
	"log"
	"time"

	sloggraylog "github.com/plinkplinkfizz/slog-graylog"
	"golang.org/x/exp/slog"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func main() {
	// docker-compose up -d
	// or
	// ncat -l 12201 -u
	gelfWriter, err := gelf.NewTCPWriter("localhost:12201")
	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}

	logger := slog.New(sloggraylog.Option{Level: slog.LevelDebug, Writer: gelfWriter}.NewGraylogHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
