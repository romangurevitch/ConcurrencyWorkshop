# 4. Finally, something works as expected

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/concurrency/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md)

```go
package concurrency

// FinallySomethingWorksAsExpected but is it?
func FinallySomethingWorksAsExpected() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			lock.Lock()
			basicCounter.Inc()
			lock.Unlock()
		}()
	}

	wg.Wait()
	return basicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="FinallySomethingWorksAsExpected$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="FinallySomethingWorksAsExpected$" -race 
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
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/l3vRfwrddpKT9ywIU/giphy.gif" width="568" alt="?"/></td>
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

[Using atomic counter](example_4_solution_atomic.md)