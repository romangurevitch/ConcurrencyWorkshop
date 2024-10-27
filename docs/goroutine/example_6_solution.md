# 6. Let's handle shutdown gracefully?

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/concurrency/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md) | [Channels](../../internal/concurrency/channel/README.md) | [Signals](../../internal/concurrency/signal/README.md)

```go
package concurrency

// NonStoppingGoRoutineWithShutdown is it good enough though?
func NonStoppingGoRoutineWithShutdown() (int, bool) {
	atomicCounter := counter.NewAtomicCounter()
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() { gracefulShutdown = true }()

		for {
			inlinePrint(atomicCounter.Inc())
		}
	}()

	<-sigs
	return atomicCounter.Count(), gracefulShutdown
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineWithShutdown$" -race 
```

<table>
<thead> 
  <tr> 
    <th colspan="3">Results?</th> 
  </tr>
</thead>
<tbody>
  <tr>
    <td>Correct result?</td>
    <td><img height="40" src="../images/no.png" width="40" alt="?"/></td> 
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/Jq824R93JsLwZCaiSL/giphy.gif" width="320" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="../images/no.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="../images/no.png" width="40" alt="?"/></td>
  </tr>
</tbody>
</table>

[Next example](example_7.md)
