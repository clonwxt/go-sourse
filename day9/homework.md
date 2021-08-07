# 作业

1. 实现一个链式栈(底层使用存储使用链表实现)

```go
type Stack struct {
	store List
}
```

2. 有一个定时器，定时器中维护了很多定时任务，每个任务都设定了一个要触发执行的时间点。定时器每过一个很小的单位时间（比如 1 秒），就扫描一遍任务，看是否有任务到达设定的执行时间。如果到达了，就拿出来执行。请进行优化。

```
2021.7.31  11:30    TaskA
2021.8.1   1:20     TaskB
2021.8.1   1:21     TaskC
2021.8.1   23:23    TaskD
```

如上所述，这样每隔 1 秒就去扫描的方法比较低效：

+ 每个任务的约定执行时间之间可能会隔很久，这样会多出很多无用的扫描。
+ 每次都去扫描整个任务列表的话，如果该表比较大，扫描时间间隔又及其短，对性能时间消耗就比较大了。