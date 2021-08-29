# gpn backend by origin dev

Решение для хакатона GPN Tech.Challenge

[Фронтенд здесь](https://github.com/alyush1n/gpn-tc-frontend)

## Стек бэкенда

* Go 1.16
* Python
* MongoDB

## Пример .env

```
MONGO_URL="mongodb://127.0.0.1:27017/?compressors=zlib&readPreference=primary&gssapiServiceName=mongodb&appname=MongoDB%20Compass&ssl=false"
PORT=8080
# ниже ссылки на парсеры
PARSE_HABR_CAREER=
PARSE_SUPPLIERS=
PARSE_RBK=
PARSE_OKVED=
```

## Запуск

развертывание сервиса производится на любой операционной системе \
требуется установленный язык Golang;

`go run app/main.go`

или скомпилировать бинарник

`go build app/main.go`
