# VK internship task
Разработка просто чат-бота на языке Go для стажировки VK.

## Основные сведения
Данная программа содержит реализацию простого чат-бота на базе VK Bots API. 

Ниже приведены основные моменты данной реализации:
 1. Реализовано на базе **LongPollServer**.
 2. **Ни одной** используемой 3rd-party библиотеки и\или зависимости, все стандартными средствами Go.
 3. Поддержка **4-х** режимов: подброс монетки, подброс кубика(-ов), получение случайного английского слова и генерация случайного числа в диапазоне.
 4. **Dockerfile** для простого запуска программы.

Структура директорий проекта выглядит следующим образом:
```
├── cmd
|   └── server
|       └── main.go                     - Запускающий файл
├── internal
|   ├── app
|   |   ├── bot.go                      - Основная логика бота
|   |   └── handlers.go                 - Вспомогательные хэндлеры для обработки каждого из событий
|   ├── config
|   |   └── config.go                   - Конфигурация бота
|   └── pkg
|       ├── api
|       |   ├── objects
|       |   |   └── message_new.go      - Object JSON-структура события "message_new" согласно спецификации VK API
|       |   └── long_poll.go            - Основные JSON-структуры для работы с VK LongPoll API
|       ├── keyboard
|       |   ├── generator.go            - Методы для генерация клавиаутры 
|       |   └── model.go                - Структуры описания генерируемой клавиатуры
|       └── operations
|           ├── coin
|           |   └── coin_message.go     - Методы для работы с операцией "Подбросить монетку"
|           ├── common
|           |   └── common_message.go   - Общие методы для отправки сообщений
|           ├── dice
|           |   └── dice_message.go     - Методы для работы с операцией "Подбросить кубик"
|           ├── number
|           |   └── number_message.go   - Методы для работы с операцией "Получить число"
|           ├── welcome
|           |   └── welcome_message.go  - Методы для отправки приветственных сообщений
|           └── word
|               ├── model.go            - Модель работы с внешним API для получения случайного слова
|               └── word_message.go     - Методы для работы с операцией "Получить слово"
├── .gitignore
├── Dockerfile
├── go.mod
├── README.md
```

## Установка и запуск
В файле cmd/server/main.go нужно указать **token** и **groupID**. Оба этих параметра могут быть получены из настроек сообщества.

Программа запускалась под macOS 13.2.1 с версией Go 1.19. Изменение любых параметров не гарантирует успешный запуск.

Проект содержит Dockerfile, поэтому рекомендуется запускать проект через него. Пример запуска:

 1. Создание образа
 ```bash
 $ sudo docker build -t vk-bot .
 ```
 ```bash
[+] Building 2.1s (8/8) FINISHED                                                                                                                   
 => [internal] load build definition from Dockerfile                                                           0.0s
 => => transferring dockerfile: 36B                                                                            0.0s
 => [internal] load .dockerignore                                                                              0.0s
 => => transferring context: 2B                                                                                0.0s
 => [internal] load metadata for docker.io/library/golang:1.19                                                 1.9s
 => [internal] load build context                                                                              0.0s
 => => transferring context: 16.78kB                                                                           0.0s
 => [1/3] FROM docker.io/library/golang:1.19@sha256:86af5649                                                   0.0s
 => CACHED [2/3] WORKDIR /app                                                                                  0.0s
 => [3/3] COPY . ./                                                                                            0.0s
 => exporting to image                                                                                         0.0s
 => => exporting layers                                                                                        0.0s
 => => writing image sha256:09bf3ee1                                                                           0.0s
 => => naming to docker.io/library/vk-bot                                                                      0.0s
 ```
 2. Запуск в контейнере
 ```bash
 $ sudo docker run --rm -d vk-bot
 ```
 3. Пример логов через интерфейс Docker Desktop
<img width="434" alt="photo" src="https://user-images.githubusercontent.com/24461208/236918405-2c4f0296-4eb3-43c0-b44f-78e20907f3fd.png">


## Тесты
... Ожидается ...
