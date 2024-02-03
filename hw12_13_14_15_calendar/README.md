#### Результатом выполнения следующих домашних заданий является сервис «Календарь»:
- [Домашнее задание №12 «Заготовка сервиса Календарь»](./docs/12_README.md)
- [Домашнее задание №13 «Внешние API от Календаря»](./docs/13_README.md)
- [Домашнее задание №14 «Кроликизация Календаря»](./docs/14_README.md)
- [Домашнее задание №15 «Докеризация и интеграционное тестирование Календаря»](./docs/15_README.md)

#### Ветки при выполнении
- `hw12_calendar` (от `master`) -> Merge Request в `master`
- `hw13_calendar` (от `hw12_calendar`) -> Merge Request в `hw12_calendar` (если уже вмержена, то в `master`)
- `hw14_calendar` (от `hw13_calendar`) -> Merge Request в `hw13_calendar` (если уже вмержена, то в `master`)
- `hw15_calendar` (от `hw14_calendar`) -> Merge Request в `hw14_calendar` (если уже вмержена, то в `master`)

**Домашнее задание не принимается, если не принято ДЗ, предшедствующее ему.**

# Useful commands
Install `sqlboiler`:
```azure
go install github.com/volatiletech/sqlboiler/v4@latest
```
Install the postgres driver:
```azure
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```
Install the following packages in the project:
```azure
go get github.com/lib/pq
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/null/v8
```
Source: [https://thedevelopercafe.com/articles/sql-in-go-with-sqlboiler-ac8efc4c5cb8](https://thedevelopercafe.com/articles/sql-in-go-with-sqlboiler-ac8efc4c5cb8)

Generate models with `Makefile`:
```azure
make models
```