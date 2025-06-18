# Базовый образ
FROM golang:1.14.4-buster
# Папка приложения
ARG APP_DIR=app
# Копирование файлов
COPY . /go/tmp/src/${APP_NAME}
# Рабочая директория
WORKDIR /go/tmp/src/${APP_NAME}