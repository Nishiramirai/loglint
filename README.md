# LogLint

Линтер для проверки лог-записей (slog, zap) в проектах на Go. Написан как плагин для `golangci-lint`.

## Правила
Линтер проверяет вызовы `log/slog` и `go.uber.org/zap` на соответствие следующим правилам:
1. Сообщение должно начинаться со строчной буквы.
2. Сообщение должно быть только на английском языке.
3. Сообщение не должно содержать спецсимволов или эмодзи (разрешены только латиница, цифры и пробелы).
4. Сообщение не должно содержать потенциально чувствительные данные (ключевые слова `password:`, `api_key=`, `token:` и т.д.).

## Сборка и установка (как кастомный плагин)

Для работы с `golangci-lint` используется Module Plugin System. Вам понадобится утилита `custom-gcl` для сборки кастомного бинарника линтера с нашим плагином.

1. Установите `custom-gcl`:
``` bash
go install github.com/golangci/golangci-lint/cmd/custom-gcl@latest
```

2. Создайте файл `.custom-gcl.yml` в корне проекта, где хотите использовать линтер:
```
    version: v1.56.2
    plugins:
      - module: 'github.com/твое_имя/loglint'
        import: 'github.com/твое_имя/loglint/plugin'
        version: main
```

3. Соберите кастомный бинарник:
``` bash
custom-gcl build .custom-gcl.yml
```

В текущей директории появится файл `custom-gcl` (или `custom-gcl.exe`).

## Использование и конфигурация

Создайте файл `.golangci.yml` со следующим содержимым:

    linters-settings:
      custom:
        loglint:
          type: module
          description: Linter for log messages

    linters:
      disable-all: true
      enable:
        - custom

Запуск линтера:
``` bash
./custom-gcl run --config .golangci.yml ./...
```

## Запуск без golangci-lint (Standalone)
Вы можете использовать линтер как самостоятельное CLI-приложение без сборки кастомного golangci-lint:
``` bash
go build -o loglint cmd/loglint/main.go ./loglint ./...
```

## Тестирование
В проекте реализованы unit-тесты для правил валидации и end-to-end тесты с использованием AST (analysistest).

    go test ./...