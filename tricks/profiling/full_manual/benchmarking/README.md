Для запуска бенча использем 
```shell
go test -bench=.

goos: darwin # информация о машине 
goarch: arm64
pkg: github.com/triumphpc/go-study/algorithms/leatcode/3_two-sum
cpu: Apple M1 Pro

BenchmarkTwoSum/small_slice-10 (это количество виртуальных ядер)        	74292668	 (сколько раз запустилась функция)        15.12 ns/op (сколько времени в среднем заняло)
BenchmarkTwoSum/medium_slice-10       	29538279	        45.04 ns/op
BenchmarkTwoSum/large_slice-10        	    4050	    255877 ns/op
BenchmarkTwoSumHashTable/small_slice-10         	23230920	        51.81 ns/op
BenchmarkTwoSumHashTable/medium_slice-10        	 3561882	       334.1 ns/op
BenchmarkTwoSumHashTable/large_slice-10         	   25123	     50787 ns/op
BenchmarkTwoSumOnePassHashTable/small_slice-10  	37688047	        31.82 ns/op
BenchmarkTwoSumOnePassHashTable/medium_slice-10 	 3552192	       324.7 ns/op
BenchmarkTwoSumOnePassHashTable/large_slice-10  	   23500	     53778 ns/op
BenchmarkTwoMy/small_slice-10                   	38387254	        37.73 ns/op
BenchmarkTwoMy/medium_slice-10                  	 3578193	       373.0 ns/op
BenchmarkTwoMy/large_slice-10                   	   22423	     51632 ns/op
PASS
ok  	github.com/triumphpc/go-study/algorithms/leatcode/3_two-sum	17.866s
```

Посмотрим какие флаги есть в функции 
```shell
go help testflag
```

Запуск специфической функции 
```shell
go test -bench='BenchmarkTwoSumWithBruteForce'
```

Запустить по два раза и по 2 секунды 
```shell
go test -bench='.' -count=2 -benchtime='2s'
```




