# СТАДИЯ 1: Builder
# Используем образ с нужной версией Go
FROM golang:1.23-alpine AS builder

WORKDIR /app

# ✅ Оптимизация: Сначала копируем только файлы зависимостей
COPY go.mod go.sum ./

# ✅ Скачиваем зависимости. Этот слой кэшируется, если go.mod/go.sum не менялись.
RUN go mod download

# Теперь копируем остальной код
COPY . .


# Собираем приложение
# CGO_ENABLED=0 для статической сборки, что делает финальный образ чище и меньше
# -o task-manager - имя выходного файла
# ./cmd/main.go - путь к твоему main файлу (подразумеваем, что он будет там)
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o task-manager ./cmd/main.go


# СТАДИЯ 2: Final Image
# Используем минималистичный образ (alpine), чтобы образ был маленький
FROM alpine:latest

# Устанавливаем сертификаты и таймзоны для корректной работы
RUN apk --no-cache add ca-certificates tzdata

# Копируем только скомпилированный бинарник из стадии builder
COPY --from=builder /app/task-manager /usr/local/bin/task-manager
COPY migration ./migration
# Порт, который будет слушать приложение (берется из .env, но указываем для документации)
EXPOSE 8080

# Точка входа в контейнер
ENTRYPOINT [ "/usr/local/bin/task-manager" ]