package core

import (
	"fmt"
	"log/slog"
	"os"
)

/*
SetupLogger настраивает и возвращает логгер в зависимости от среды выполнения (env).

Логгер конфигурируется по-разному для сред Local, Dev и Prod:

- В среде Local используется текстовый обработчик с уровнем отладки (Debug) и добавлением источника вызова.

- В среде Dev используется JSON-обработчик с уровнем отладки (Debug) и добавлением источника вызова.

- В среде Prod используется JSON-обработчик с уровнем информации (Info) и добавлением источника вызова.

Если переданное значение среды не поддерживается, функция вызывает панику.
*/
func SetupLogger(env EnvType) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case Local:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}),
		)
	case Dev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}),
		)
	case Prod:

		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelInfo,
				AddSource: true,
			}),
		)
	default:
		panic(fmt.Sprintf("env: %s not in [%s, %s, %s]", env, Local, Dev, Prod))
	}

	return logger
}
