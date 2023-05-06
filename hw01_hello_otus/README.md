## Домашнее задание №1 «Hello, OTUS!»

Необходимо написать программу, печатающую в стандартный вывод перевернутую фразу
```
Hello, OTUS!
```

Для переворота строки следует воспользоваться возможностями
[golang.org/x/example/stringutil](https://github.com/golang/example/tree/master/stringutil).

Кроме этого необходимо исправить **go.mod** так, чтобы для данного модуля работала
команда `go get`, а полученный **go.sum** закоммитить.

### Критерии оценки
- Пайплайн зелёный - 4 балла
- Используется `stringutil` - 4 балла
- Понятность и чистота кода - до 2 баллов

#### Зачёт от 7 баллов

### Подсказки
- `Reverse`


## Commands

 - Пишем код:
```
gofmt -w main.go
go run main.go
```

 - Проверияем что все ок:
```
golangci-lint run .
```

 - ` --- когда не срабоатет ---`
```
go install github.com/daixiang0/gci@latest
gci write --skip-generated  .
# или
golangci-lint run --fix
```

 - После исправлений перепроверям:
```
go mod tidy
gofmt -w main.go
go run main.go
```

