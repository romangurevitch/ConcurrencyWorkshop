# 7. Let's handle shutdown gracefully, for real this time!

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/buildingblocks/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md) | [Channels](../../internal/buildingblocks/channel/README.md) | [Signals](../../internal/buildingblocks/signal/README.md)

```go
package concurrency

// NonStoppingGoRoutineCorrectShutdown yes?
func NonStoppingGoRoutineCorrectShutdown() (int, bool) {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}
	gracefulShutdown := false

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { gracefulShutdown = true }()

		for {
			select {
			case <-sigs:
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
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineCorrectShutdown$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutineCorrectShutdown$" -race 
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
    <td><img height="40" src="../images/yes.png" width="40" alt="?"/></td>
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/3oxRmD9a5pLTOOLigM/giphy.gif" width="320" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="../images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="../images/yes.png" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Next example](example_8.md)

