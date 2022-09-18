## Создать файл для миграции
```sh
migrate create -ext sql -dir db/migrations {FILE_NAME}
```
## Применить миграции пример
```sh
migrate -database 'postgres://{USERNAME}:{PASSWORD}@{DBIP}:{DBPORT}/{DBNAME}?sslmode=disable&search_path={SCHEMA_NAME}' -path db/migrations/ up
```

## Применить миграции локально
```sh
migrate -database 'postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable&search_path=gameserver' -path db/migrations/ up
```