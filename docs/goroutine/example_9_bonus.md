# 10. Bonus: let's make a tiny change

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/buildingblocks/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md) | [Channels](../../internal/buildingblocks/channel/README.md) | [Signals](../../internal/buildingblocks/signal/README.md) | [Context](../../internal/buildingblocks/context/README.md)

```go
package concurrency

// NonStoppingGoRoutineContextBonus use context with tiny change
func NonStoppingGoRoutineContextBonus(ctx context.Context) (int, bool) {
	atomicCounter := counter.NewAtomicCounter()

	ctx, cancelFunc := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()
	wg := sync.WaitGroup{}
	gracefulShutdown := false

	wg.Add(1)
	go func() {
		defer func() { gracefulShutdown = true }() // ↑ ↓ switch 
		defer wg.Done()

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
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContextBonus$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineContextBonus$" -race 
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
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/U1TgwOffGUqxQYClV1/giphy.gif" alt="?"/></td>
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

[Questions?](questions.md)