# 8. Adding context

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/concurrency/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md) | [Channels](../../internal/concurrency/channel/README.md) | [Signals](../../internal/concurrency/signal/README.md) | [Context](../../internal/concurrency/context/README.md)

```go
package concurrency

// NonStoppingGoRoutineContext use context
func NonStoppingGoRoutineContext(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	gracefulShutdown := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()
		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down goroutine", "reason", ctx.Err())
				return
			case reason := <-sigs:
				slog.Info("shutting down goroutine", "reason", reason)
				return
			default:
				inlinePrint(atomicCounter.Inc())
			}
		}
	}()

	wg.Wait()
	return atomicCounter.Count(), gracefulShutdown
}
```

```bash
 clear; go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContext$" 
```

```bash
 clear; go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContext$" -race 
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
    <td><img height="40" src="../images/question.svg" width="40" alt="?"/></td>
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/f3txRVhVh0MOO5vZkB/giphy.gif" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="../images/question.svg" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="../images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Solution](example_8_solution.md)