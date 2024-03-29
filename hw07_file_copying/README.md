## Домашнее задание №7 «Утилита для копирования файлов»
Необходимо реализовать утилиту копирования файлов (упрощенный аналог `dd`).

Тулза должна принимать следующие аргументы:
* путь к исходному файлу (`-from`);
* путь к копии (`-to`);
* отступ в источнике (`-offset`), по умолчанию - 0;
* количество копируемых байт (`-limit`), по умолчанию - 0 (весь файл из `-from`).

Особенности:
* offset больше, чем размер файла - невалидная ситуация;
* limit больше, чем размер файла - валидная ситуация, копируется исходный файл до его EOF;
* программа может НЕ обрабатывать файлы, у которых неизвестна длина (например, /dev/urandom);

Также необходимо выводить в консоль прогресс копирования в процентах (%),
допускается использовать для этого стороннюю библиотеку.

Юнит-тесты могут использовать файлы из `testdata` (разрешено добавить свои, но запрещено удалять имеющиеся)
и должны чистить за собой создаваемые файлы (или работать в `/tmp`).

При необходимости можно выделять дополнительные функции / ошибки.

**(*) Дополнительное задание: реализовать прогресс-бар самостоятельно.**

### Критерии оценки
- Пайплайн зелёный - 4 балла
- Добавлены юнит-тесты - до 4 баллов
- Понятность и чистота кода - до 2 баллов
- Дополнительное задание на баллы не влияет

#### Зачёт от 7 баллов

### Подсказки
- `github.com/cheggaaa/pb`
- `os.OpenFile`, `os.Create`, `os.FileMode`
- `io.CopyN`
- `os.CreateTemp`

### useful links
 - [https://stackoverflow.com/a/50312711](https://stackoverflow.com/a/50312711)
 - [https://pkg.go.dev/os#File.Seek](https://pkg.go.dev/os#File.Seek)
 - [https://pkg.go.dev/io#CopyN](https://pkg.go.dev/io#CopyN)

### useful commands
```azure
go get github.com/cheggaaa/pb/v3
go get github.com/stretchr/testify/require
```