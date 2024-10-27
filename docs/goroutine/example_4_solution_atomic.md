# 4. Finally, something works as expected - using atomic counter

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/concurrency/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md)

```go
package concurrency

// FinallySomethingWorksAsExpectedAtomicCounter but is it?
func FinallySomethingWorksAsExpectedAtomicCounter() int {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicCounter.Inc()
		}()
	}

	wg.Wait()
	return atomicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="FinallySomethingWorksAsExpectedAtomicCounter$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="FinallySomethingWorksAsExpectedAtomicCounter$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/f9Rrghj6TDckb5nZZR/giphy.gif" width="360" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="../images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="../images/question.svg" width="40" alt="?"/></td>
  </tr>
</tbody>
</table>

[Next example](example_5.md)