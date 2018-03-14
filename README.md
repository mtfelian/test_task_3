# Test task 3

This is a test task for `Imhio Ltd`.

## Deploying the app

- First of all, you should install and configure Go and PostgreSQL.
- Clone this repo.
- After installing, create the database, from `psql` console:

```
CREATE DATABASE "test_task_3"
```

- Edit DSN in `goose/dbconf.yml` according to your DB auth configuration:

- Next install the migration tool and apply the migrations:

```
go get bitbucket.org/liamstask/goose/cmd/goose
goose -path=goose up
```

### Testing

- Ensure that DSN inside func `PostgresDSNTests()` in
`./storage/postgresql.go` returns the right value.
- Launch `go test ./...` or `ginkgo -r` (if you installed Ginkgo)

Output:
```
felian@felian-VirtualBox:~/go_code/src/github.com/mtfelian/test_task_3$ ginkgo -r
[1521068786] Main Suite - 6/6 specs connecting to DB: host=localhost port=5432 user=postgres dbname=test_task_3 sslmode=disable client_encoding=utf8
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /param                    --> github.com/mtfelian/test_task_3.getParam (3 handlers)
[GIN] 2018/03/15 - 02:06:28 | 200 |    4.386215ms |       127.0.0.1 |  POST     /param
•
(sql: database is closed)
[2018-03-15 02:06:28]
[GIN] 2018/03/15 - 02:06:28 | 500 |    1.803276ms |       127.0.0.1 |  POST     /param
connecting to DB: host=localhost port=5432 user=postgres dbname=test_task_3 sslmode=disable client_encoding=utf8
•[GIN] 2018/03/15 - 02:06:28 | 200 |    4.582904ms |       127.0.0.1 |  POST     /param
•[GIN] 2018/03/15 - 02:06:28 | 422 |      36.889µs |       127.0.0.1 |  POST     /param
•[GIN] 2018/03/15 - 02:06:28 | 422 |      60.184µs |       127.0.0.1 |  POST     /param
•• SUCCESS! 27.573609ms PASS
[1521068786] Config Suite - 1/1 specs • SUCCESS! 921.928µs PASS

Ginkgo ran 2 suites in 1.941966404s
Test Suite Passed
```

### Running

```
go build && ./test_task_3 --port=3000 --db='host=localhost port=5432 user=postgres dbname=test_task_3 sslmode=disable client_encoding=utf8'
```

#### Params:

- `port` to specify HTTP server port
- `db` to configure the DSN

### API

Standard error:
```
{
  "code": 1,
  "error": "text"
}
```

Codes:
```
1 - failed to parse input data
2 - validation on input data failed
3 - storage access error
```

#### POST /param

Get a log entry with given data in body.

Request body JSON:
```
{
  "Type": "Test.vpn",
  "Data": "Rabbit.log"
}
```

Response body contains a string (not a JSON Content-Type) with JSON
object inside.

##### Response HTTP:
- 200 - OK
- 422 - Validation/input error
- 500 - Storage error

###### OK:
Empty body.

###### Error:
Standard error.

## Source text

```
Разработать сервис-конфигуратор.

Требования к сервису:
- После запуска сервис слушает запросы на определённом порте;
- В качестве запроса сервис получает JSON-строку, представляющую объект, содержаший название конфигурации и название блока параметров;
- Название блока параметров определяет набор и типы параметров, а название конфигурации определяет конкретные значения;
- В ответ сервис отправляет строку с JSON-объектом, содержащим запрошенные параметры.

Требования к реализации:
- Параметры сервис получает из базы данных;
- База данных заполняется параметрами до запуска сервиса;
- В качестве СУБД желательно (но не обязательно) выбрать PostgreSQL.

Требования к разработке:
- Разработка ведётся от тестов (выбор TDD/BDD/DDD и модуля тестирования - на усмотрение исполнителя);
- Нетривиальные участки кода коментируются;
- Для взаимодействия с СУБД использовать существующий модуль ORM (любой);
- Для первоначальной вставки данных в СУБД использовать модуль миграции (любой).

Требования к результату:
- Предоставить миграции (вверх и вниз) для создания хранилица конфигураций в СУБД и посева тестовых данных.
- Предоставить код, организованный в репозиторий на Github или Bitbucket(+) или в архиве(-).
- Покрытие кода тестами должно быть не менее 70%

Пример запроса:
{
"Type": "Develop.mr_robot",
"Data": "Database.processing"
}

Пример ответа:
{
"host": "localhost",
"port": "5432",
"database": "devdb",
"user": "mr_robot",
"password": "secret",
"schema": "public"
}

Запрос:
{
"Type": "Test.vpn",
"Data": "Rabbit.log"
}

Ответ:
{
"host": "10.0.5.42",
"port": "5671",
"virtualhost": "/",
"user": "guest",
"password": "guest"
}

```