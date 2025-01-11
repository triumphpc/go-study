```shell
go test -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'
go tool pprof mem.prof
File: cpu.test
Type: alloc_space
Time: Jan 11, 2025 at 2:48pm (MSK)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
```

Посомтреть какая кода сколько памяти сожрала 
```shell
(pprof) top10
Showing nodes accounting for 14724.13MB, 100% of 14728.51MB total
Dropped 30 nodes (cum <= 73.64MB)
      flat  flat%   sum%        cum   cum%
 7786.20MB 52.86% 52.86%  7786.20MB 52.86%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.twoSumOnePassHashTable (inline)
 2972.36MB 20.18% 73.05%  2972.36MB 20.18%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.twoSumHashTable (inline)
 1671.53MB 11.35% 84.39%  1671.53MB 11.35%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.BenchmarkTwoSum.func1
 1232.52MB  8.37% 92.76%  5120.89MB 34.77%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.BenchmarkTwoMy.func1
  630.01MB  4.28% 97.04%  4527.83MB 30.74%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.BenchmarkTwoSumOnePassHashTable.func1
  431.51MB  2.93%   100%  3403.87MB 23.11%  github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.BenchmarkTwoSumHashTable.func1
         0     0%   100% 14723.62MB   100%  testing.(*B).launch
         0     0%   100% 14724.63MB   100%  testing.(*B).runN
```

```shell
(pprof) list twoSum
Total: 14.38GB
ROUTINE ======================== github.com/triumphpc/go-study/tricks/profiling/full_manual/cpu.twoSumHashTable in /Users/s.vrulin/Devel/go-study/tricks/profiling/full_manual/cpu/main_test.go
    2.90GB     2.90GB (flat, cum) 20.18% of Total
         .          .    105:
         .          .    106:func twoSumHashTable(nums []int, target int) []int {
         .          .    107:	m := make(map[int]int)
    2.90GB     2.90GB    108:	for i, num := range nums {
         .          .    109:		m[num] = i
         .          .    110:	}
         .          .    111:	for i, num := range nums {
         .          .    112:		complement := target - num
         .          .    113:		if j, ok := m[complement]; ok && j != i {
```

![img.png](img.png)