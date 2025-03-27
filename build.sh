#!/bin/bash

# Создаем директорию для сборки если её нет
mkdir -p build

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
    
    # Запускаем сборку
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name"
    
    if [ $? -ne 0 ]; then
        echo "Ошибка при сборке для $GOOS/$GOARCH"
        exit 1
    fi
done

echo "Сборка завершена успешно!" 