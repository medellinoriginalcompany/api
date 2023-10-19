package config

import (
	"log/slog"
	"os"
	"time"
)

func LoggerHandler() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Selecionar o tipo de log como debug, info, warn, error
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "date" // Renomear para date
				a.Value = slog.StringValue(time.Now().Format("2006.01.02 03:04:05")) // Formatar a data
			}
			return a
		},
	}).WithAttrs([]slog.Attr{ // Adicionar atributos globais para cada log
		slog.String("app", "Medellin Original Company"),
		slog.String("version", "0.1.0"),
	},
	)

	logger := slog.New(logHandler)
	slog.SetDefault(logger) // Configurar slog

	logfile, err := os.Create("app.log")

	if err != nil {
		slog.Error(err.Error())
	}

	defer logfile.Close()

}