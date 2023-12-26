## Домашнее задание №10 «Оптимизация программы»

Вам дан исходный код функции `GetDomainStat(r io.Reader, domain string)`, которая:
* читает построчно из `r` пользовательские данные вида
```text
{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Brian Olson","Username":"non_quia_id","Email":"FrancesEllis@Quinu.edu","Phone":"237-75-34","Password":"cmEPhX8","Address":"Butterfield Junction 74"}
{"Id":3,"Name":"Justin Oliver Jr. Sr.","Username":"oPerez","Email":"MelissaGutierrez@Twinte.gov","Phone":"106-05-18","Password":"f00GKr9i","Address":"Oak Valley Lane 19"}
```
(осторожно, в отличие от конкретной строки файл целиком не является валидным JSON);
* подсчитывает количество email-доменов пользователей на основе домена первого уровня `domain`.

Например, для данных, представленных выше:
```text
GetDomainStat(r, "com") // {}
GetDomainStat(r, "gov") // {"browsedrive": 1, "twinte": 1}
GetDomainStat(r, "edu") // {"quinu": 1}
```

Для большего понимания см. исходный код и тесты.

**Необходимо оптимизировать программу таким образом, чтобы она проходила все тесты.**

Нельзя:
- изменять сигнатуру функции `GetDomainStat`;
- удалять или изменять существующие юнит-тесты.

Можно:
- писать любой новый необходимый код;
- удалять имеющийся лишний код (кроме функции `GetDomainStat`);
- использовать сторонние библиотеки по ускорению анмаршалинга JSON;
- добавлять юнит-тесты.

**Обратите внимание на запуск TestGetDomainStat_Time_And_Memory**
```bash
go test -v -count=1 -timeout=30s -tags bench .
```

Здесь используется билд-тэг bench, чтобы отделить обычные тесты от тестов производительности.

### Оформление пул-риквеста
В идеале к подобным пул-риквестам пишут бенчмарки и прикладывают результаты работы benchstat, чтобы сразу было видно, что стало лучше и насколько.

### Критерии оценки
- Пайплайн зелёный и нет попытки «обмануть» систему - 4 балла
- Добавлены юнит-тесты - до 3 баллов
- Понятность и чистота кода - до 3 баллов

### Частые ошибки
- Работа с сырыми байтами, нахождение позиции `"Email"` и пр. вместо ускорения анмаршалинга более поддерживаемыми и понятными средствами.

#### Зачёт от 7 баллов


## Заупск бенчмарков
```sh
go test -v -count=1 -timeout=30s -tags bench . > out.txt
```

```text
=== RUN   TestGetDomainStat_Time_And_Memory
    stats_optimization_test.go:46: time used: 337.991833ms / 300ms
    stats_optimization_test.go:47: memory used: 308Mb / 30Mb
--- PASS: TestGetDomainStat_Time_And_Memory (6.11s)
PASS
ok  	github.com/fixme_my_friend/hw10_program_optimization	6.443s
```

```sh
go test -bench=BenchmarkGetDomainStat -cpuprofile=cpu.out -memprofile=mem.out .
brew install graphviz
go tool pprof -http=":8090" hw10_program_optimization.test cpu.out
```

```text
goos: darwin
goarch: arm64
pkg: github.com/fixme_my_friend/hw10_program_optimization
BenchmarkGetDomainStat
BenchmarkGetDomainStat-8   	      16	  71786885 ns/op	136045282 B/op	 1700090 allocs/op
BenchmarkGetUsers
BenchmarkGetUsers-8        	    1880	    619661 ns/op	    6296 B/op	      66 allocs/op
BenchmarkCountDomains
BenchmarkCountDomains-8    	      14	  73049366 ns/op	136039284 B/op	 1700029 allocs/op
PASS
ok  	github.com/fixme_my_friend/hw10_program_optimization	5.244s
```