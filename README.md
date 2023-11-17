Установка:
```
make compose
```
Подождать пока контейнер `twirler` перестанет перезапускаться [ждёт кафку : )]
```
make migrate-up
```

Тесты:
```
make test-unit
```
Для интеграционных нужно БД
```
make test-integration
```