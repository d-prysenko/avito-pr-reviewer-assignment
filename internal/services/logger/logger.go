package logger

import (
	"log/slog"
	"os"
	"revass/internal/config"
	"time"
)

func SetupLogger(cfg config.Config) *slog.Logger {
	var log *slog.Logger

	if cfg.Env == config.EnvLocal {
		opts := PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
		}

		handler := opts.NewPrettyHandler(os.Stdout)

		log = slog.New(handler)
	}

	if cfg.Env == config.EnvProd {
		currentTime := time.Now()
		filename := "logs/log_" + currentTime.Format("02_01_2006__15_04_05")

		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("error opening file: " + err.Error())
		}

		log = slog.New(
			slog.NewJSONHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
