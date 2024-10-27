# 9. Notify context on signals

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/concurrency/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md) | [Channels](../../internal/concurrency/channel/README.md) | [Signals](../../internal/concurrency/signal/README.md) | [Context](../../internal/concurrency/context/README.md)

```go
package concurrency

// NonStoppingGoRoutineContextBetter use context
func NonStoppingGoRoutineContextBetter(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	ctx, cancelFunc := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()
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
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContextBetter$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContextBetter$" -race 
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
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/m1TY53YqzzTY0Efq63/giphy.gif" alt="?"/></td>
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

[Solution](example_9_solution.md)