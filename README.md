# VK Internship Task
Разработка просто чат-бота на языке Go для стажировки VK.
<details>
  <summary>Содержание</summary>
  <ol>
    <li><a href="#основные-сведения">Основные сведения</a></li>
    <li><a href="#установка-и-запуск">Установка и запуск</a></li>
    <li><a href="#тесты-и-ci">Тесты и CI</a></li>
  </ol>
</details>

## Основные сведения
Данная программа содержит реализацию простого чат-бота на базе VK Bots API. 

Ссылка на бота: _https://vk.com/club220352548_

Ниже приведены основные моменты данной реализации:
 1. Реализовано на базе **LongPollServer**.
 2. **Ни одной** используемой 3rd-party библиотеки и\или зависимости, все стандартными средствами Go.
 3. Поддержка **5-и** режимов:
  
    3.1. **Подброс монетки** — _пользователь выбирает между Орлом и Решкой._
    </br>
    
    3.2. **Подброс кубика(-ов)** — _полезно для игры, где нужно бросать игральные кости, но их нет под рукой. Можно одновременно бросить до 3-х костей._
    </br>
    
    3.3. **Получение случайного английского слова** — _бот выдает случайное английское слово. Сойдет, например, для игры в слова, alias. Можно выбрать между тремя опциями: Существительное\Прилагательное\Животное._
    </br>
    
    3.4. **Генерация случайного числа в диапазоне** — _не нуждается в пояснении._
    </br>
    
    3.5. **Создание комнаты** — _создает комнату, к которой могут подключиться другие люди, чтобы вместе вести счет игры (например, для игры в монополию, бильярд, и т. д.)._
 4. **Dockerfile** для простого запуска программы.
 5. Простой **CI** и тесты для проверки минимального функционала.

Структура директорий проекта выглядит следующим образом:
```markdown
├── .github
|   └── workflows
|       └── test.yml                    - Файл конфигурации CI
├── build
|   └── Dockerfile                      - Docker образ для запуска бота
├── cmd
|   └── server
|       └── main.go                     - Запускающий файл
├── internal
|   ├── app
|   |   ├── bot.go                      - Основная логика бота
|   |   ├── bot_test.go                 - Тесты для главной логики бота
|   |   ├── mode_handlers.go            - Методы для обработки каждого из режимов
|   |   └── update_handlers.go          - Хэндлеры для обработки каждого из событий VK API
|   ├── config
|   |   └── config.go                   - Конфигурация бота
|   └── pkg
|       ├── api
|       |   ├── objects
|       |   |   └── message_new.go      - Object JSON-структура события "message_new" согласно спецификации VK API
|       |   └── long_poll.go            - Основные JSON-структуры для работы с VK LongPoll API
|       ├── keyboard
|       |   ├── generator.go            - Методы для генерация клавиатуры 
|       |   └── model.go                - Структуры описания генерируемой клавиатуры
|       └── operations
|           ├── coin
|           |   └── vk_methods.go       - Методы для работы с VK API для операции "Подбросить монетку"
|           ├── common
|           |   └── vk_methods.go       - Общие методы для работы с VK API для отправки сообщений
|           ├── dice
|           |   └── vk_methods.go       - Методы для работы с VK API для операции "Подбросить кубик"
|           ├── number
|           |   └── vk_methods.go       - Методы для работы с VK API для операции "Получить число"
|           ├── room
|           |   ├── model.go            - Структуры для работы с комнатами для операции "Создать/Найти комнату"
|           |   ├── room_methods.go     - Методы для работы и взаимодействия с комнатами
|           |   └── vk_methods.go       - Методы для работы с VK API для операции "Создать/Найти комнату"
|           ├── welcome
|           |   └── vk_methods.go       - Методы для работы с VK API для отправки приветственных сообщений
|           └── word
|               ├── model.go            - Модель работы с внешним API для получения случайного слова
|               └── vk_methods.go       - Методы для работы с VK API для операции "Получить слово"
├── tools
|   └── setup_env.go                    - Вспомогательная утилита для установки переменных окружения из .env файла
├── .env
├── .gitignore
├── go.mod
└── README.md
```

## Установка и запуск
В файле **".env"**, который имеет следующий вид:
 ```env
TOKEN=YOUR_API_TOKEN
GROUP_ID=YOUR_GROUP_ID
 ```
Вместо **"YOUR_API_TOKEN"** и **"YOUR_GROUP_ID"** нужно указать значения **Ключа доступа** и **ID группы**, от лица которой бот будет отправлять сообщения. Оба этих параметра могут быть получены из настроек сообщества.

Проект содержит **build/Dockerfile**, поэтому рекомендуется запускать проект через него. Пример запуска:

 1. Создание образа (запускается в корне проекта)
 ```bash
 $ docker build -f build/Dockerfile -t vk-bot .
 ```
 ```bash
[+] Building 101.0s (16/16) FINISHED                                                                                                               
 => [internal] load build definition from Dockerfile                                0.0s
 => => transferring dockerfile: 37B                                                 0.0s
 => [internal] load .dockerignore                                                   0.0s
 => => transferring context: 2B                                                     0.0s
 => [internal] load metadata for docker.io/library/golang:alpine                    2.0s
 => [internal] load metadata for docker.io/library/alpine:latest                    2.1s
 => [auth] library/golang:pull token for registry-1.docker.io                       0.0s
 => [auth] library/alpine:pull token for registry-1.docker.io                       0.0s
 => [stage-1 1/4] FROM docker.io/library/alpine@sha256:02bb6f                       0.0s
 => [builder 1/4] FROM docker.io/library/golang:alpine@sha256:913de9                0.0s
 => [internal] load build context                                                   0.1s
 => => transferring context: 30.90kB                                                0.0s
 => CACHED [builder 2/4] WORKDIR /app                                               0.0s
 => [builder 3/4] COPY . .                                                          0.1s
 => [builder 4/4] RUN go build -o bot cmd/server/main.go                            98.7s
 => CACHED [stage-1 2/4] WORKDIR /app                                               0.0s
 => CACHED [stage-1 3/4] COPY --from=builder /app/bot .                             0.0s
 => CACHED [stage-1 4/4] COPY --from=builder /app/.env .                            0.0s
 => exporting to image                                                              0.0s
 => => exporting layers                                                             0.0s
 => => writing image sha256:87e804                                                  0.0s
 => => naming to docker.io/library/vk-bot                                           0.0s
 ```
 2. Запуск в контейнере
 ```bash
 $ docker run --rm -d vk-bot
 ```
 3. Пример логов через интерфейс Docker Desktop
<img width="434" alt="photo" src="https://user-images.githubusercontent.com/24461208/236918405-2c4f0296-4eb3-43c0-b44f-78e20907f3fd.png">

## Тесты и CI

Проект содержит один файл **(internal/app/bot_test.go)**, который тестирует главный и минимальный функционал приложения.

В проекте был настроен простой CI для запуска кода на тестах.
Файл конфигурации может быть найден по следующему пути: **.github/workflows/test.yml**
и имеет следующий вид:
```yml
name: test

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.19

      - name: Set up environment variables
        run: |
           echo "TOKEN=${{ secrets.API_KEY }}" >> .env
           echo "GROUP_ID=${{ secrets.GROUP_ID }}" >> .env

      - name: Test
        run: go test -v VK-bot/internal/app
 ```
**"secrets.API_KEY"** и **"secrets.GROUP_ID"** это существующие данные для запуска бота,
которые скрыты от публичного доступа. Поэтому, при каждом push или pull-request код запускается на тестах с этими данными.

<p align="right">(<a href="#основные-сведения">К началу</a>)</p>
