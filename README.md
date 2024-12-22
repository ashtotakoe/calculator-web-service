# Калькулятор арифметических выражений

Веб-сервис для вычисления арифметических выражений. Сервис принимает математическое выражение через HTTP запрос и возвращает результат вычисления.

## Запуск проекта

```bash
# Клонируем репозиторий
git clone https://github.com/ashtotakoe/calculator-web-service
cd calculator-web-service
# Запуск сервера
go run cmd/main.go [порт]
```

Для отображения детальных ошибок, нужно после порта добавить флаг **detailed-validation**

```bash
  go run cmd/main.go [порт] detailed-validation
```

## Примеры использования

### 1. Успешное вычисление выражения

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'
```

Ответ (HTTP 200):

```json
{
  "result": "6"
}
```

### 2. Некорректное выражение

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*a"
}'
```

Ответ (HTTP 422):

```json
{
  "error": "Expression is not valid" // или более детальная ошибка (флаг detailed-validation)
}
```

### 3. Также может вернуться internal server error

Ответ (HTTP 500):

```json
{
  "error": "Internal server error"
}
```
