## Развертывание и запуск проекта

1. Собрать образ проекта `docker build -t egame -f docker/app/Dockerfile .`
2. Настроить переменные окружения в 'docker-compose.yml'
   - APP_CHAN_SIZE - размер буфера канала
   - APP_CHUNK_SIZE - размер чанков для записи в DB
3. Запустить локально, выполнив `docker-compose up egame`

### Настройка базы данных

Необходимо применить содержимое файла `app/infrastructure/storage/schema.sql`

## Запуск нагрузочного тестирования

1. Перейти в директорию `cd ./test`
2. Указать в файле `token.txt` свой токен от Яндекс.Танк `https://overload.yandex.net`
3. Запустить сценарий тестирования
    - Для Windows `docker run --rm -v %cd%:/var/loadtest -v  C:\\%HOMEPATH%\.ssh\:/root/.ssh -it direvius/yandex-tank`
    - Для Unis `docker run --rm -v $(pwd):/var/loadtest -v $HOME/.ssh:/root/.ssh -it direvius/yandex-tank`