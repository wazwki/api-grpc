## Использование приложения

#### — Запускает линтер  
```sh
make lint
```

#### — Запускает тесты  
```sh
make test
```

#### — Собирает Docker-образы  
```sh
make build
```

#### — Запускает приложение в фоне  
```sh
make up
```

#### — Выполняет все шаги (lint, test, build, up)  
```sh
make run
```

#### — Выполняет все шаги и запускает приложение с отключенным jwt-токенами и включенным swagger api
```sh
make run DEBUG=true
```

#### — Останавливает приложение
```sh
make stop
```

#### — Запускает приложение в debug режиме с отключенным jwt-токенами, включенным swagger api и возможностью изменения кода в запущенном приложении(docker-compose) без пересборки
```sh
make debug
```

#### — Останавливает приложение в debug режиме
```sh
make debug-stop
```

#### — Генерирует swagger-документацию
```sh
make swagger
```

#### — Генерирует swagger-документацию с заданным url
```sh
make swagger SWAGGER_URL="swagger.json"
```

#### — Генерирует proto файлы
```sh
make proto
```

#### Если golangci-lint или docker-compose не установлены, их нужно поставить.  

#### http://localhost:8081/swagger-ui/ - адрес swagger, для использования запускать приложение в debug режиме  

---