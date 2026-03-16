package example

import "log/slog"

func example() {
	password := "12345"
	apiKey := "secret-key"
	token := "jwt-token"

	// Лог-сообщения должны начинаться со строчной буквы

	slog.Info("Starting server on port 8080")   // want "log message must start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log message must start with a lowercase letter"

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	// Только на английском языке ---

	slog.Info("запуск сервера")                    // want "log message must be in English only"
	slog.Error("ошибка подключения к базе данных") // want "log message must be in English only"

	slog.Info("starting server")

	// Не должны содержать спецсимволы или эмодзи

	slog.Info("server started! 🚀")                // want "log message must not contain special characters or emojis"
	slog.Error("connection failed!!!")            // want "log message must not contain special characters or emojis"
	slog.Warn("warning: something went wrong...") // want "log message must not contain special characters or emojis"

	slog.Info("server started")
	slog.Warn("something went wrong")

	// Не должны содержать чувствительные данные

	slog.Info("user password: " + password) // want "log message must not contain special characters or emojis" "log message must not contain sensitive data"
	slog.Debug("api_key=" + apiKey)         // want "log message must not contain special characters or emojis" "log message must not contain sensitive data"
	slog.Info("token: " + token)            // want "log message must not contain special characters or emojis" "log message must not contain sensitive data"

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")

}
