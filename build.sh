#!/bin/bash

# Создаем директорию для сборки если её нет
mkdir -p build

# Определяем текущую ОС для правильной компиляции
HOST_OS=$(go env GOOS)
HOST_ARCH=$(go env GOARCH)

# Загружаем зависимости
echo "Загрузка зависимостей..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "Ошибка при загрузке зависимостей"
    exit 1
fi

# Список платформ для сборки
PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/386")

# Перебираем все платформы и собираем под каждую
for platform in "${PLATFORMS[@]}"
do
    # Разделяем строку платформы на OS и ARCH
    IFS='/' read -r -a array <<< "$platform"
    GOOS="${array[0]}"
    GOARCH="${array[1]}"
    
    # Формируем имя исполняемого файла
    output_name="build/bookmark-api-$GOOS-$GOARCH"
    if [ $GOOS = "windows" ]; then
        output_name+=".exe"
    fi

    # Выводим информацию о текущей сборке
    echo "Сборка для $GOOS/$GOARCH"
    
    # Для всех платформ используем чистый Go с тегами для modernc.org/sqlite
    env CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -tags "moderncsqlite,libsqlite3" -o "$output_name"
    
    if [ $? -ne 0 ]; then
        echo "Ошибка при сборке для $GOOS/$GOARCH"
        exit 1
    fi
done

echo "Сборка завершена успешно!"
