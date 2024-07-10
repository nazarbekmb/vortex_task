# Документация проекта 
## Описание проекта

Проект "Statistics Collection" представляет собой сервис для сбора и хранения статистических данных о заказах.
### Используемые технологии

- Язык программирования: Go (Golang)
- База данных: PostgreSQL
- HTTP сервер: стандартная библиотека net/http
- Миграции базы данных: golang-migrate

## Структура проекта

Проект организован в следующих директориях и модулях:

- config: модуль для загрузки конфигураций из .env файла.
- database: модуль для инициализации и взаимодействия с PostgreSQL базой данных, включая функции для сохранения и получения данных.
- handlers: обработчики HTTP запросов для обработки заказов и истории заказов.
- migrations: миграции базы данных PostgreSQL.
- models: структуры данных для описания заказов, истории заказов и элементов Order Book.
- routers: инициализация маршрутов HTTP API.

## Конфигурация

Конфигурационные параметры загружаются из .env файла, если он доступен. В противном случае используются значения по умолчанию.
Использование API
Запросы HTTP API

- GET /api/get_order_book: получение Order Book для указанной биржи и валютной пары.
- POST /api/save_order_book: сохранение Order Book в базу данных.
- GET /api/get_order_history: получение истории заказов для указанного клиента.
- POST /api/save_order_history: сохранение заказа в историю в базу данных.

## Параметры запросов

- Для /api/get_order_book и /api/get_order_history: необходимо указать параметры exchange_name и pair.
- Для /api/save_order_book и /api/save_order_history: передается JSON с данными заказа или Order Book.

## Запуск приложения

Приложение запускается на порту 8080. Для запуска необходимо выполнить следующие шаги:

1. Установить необходимые зависимости: Go, PostgreSQL.
2. Настроить .env файл с параметрами DATABASE_URL и SERVER_PORT.
3. Применить миграции базы данных с помощью команды migrate -database <DATABASE_URL> -path migrations up.
4. Запустить приложение: выполнить go run cmd/main.go.

## Примеры использования
Пример запроса на получение Order Book:

```
GET /api/get_order_book?exchange_name=exchange1&pair=BTC_USD
```

Пример запроса на сохранение Order Book:

```
POST /api/save_order_book
{
  "exchange": "exchange1",
  "pair": "BTC_USD",
  "asks": [
    {"price": 35000, "base_qty": 2},
    {"price": 35500, "base_qty": 3}
  ],
  "bids": [
    {"price": 34000, "base_qty": 1},
    {"price": 34500, "base_qty": 2}
  ]
}
```

Пример запроса на получение Order History

```
GET /api/get_order_history?client_name=John Doe&exchange_name=Binance&label=Order123&pair=BTC/USD
```

Пример запроса на сохранение Order History

```
POST /api/save_order_history
{
    "client_name": "John Doe",
    "exchange_name": "Binance",
    "label": "Order123",
    "pair": "BTC/USD",
    "side": "buy",
    "type": "limit",
    "base_qty": 1.5,
    "price": 45000,
    "algorithm_name_placed": "Algo1",
    "lowest_sell_prc": 44000,
    "highest_buy_prc": 46000,
    "commission_quote_qty": 10,
    "time_placed": "2024-07-10T10:00:00Z"
}
```
## Ошибки и обработка исключений

При возникновении ошибок сервер возвращает соответствующий HTTP статус и сообщение об ошибке в формате JSON.