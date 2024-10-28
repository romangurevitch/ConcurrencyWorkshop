# 5. Non-stopping go routines

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/buildingblocks/sync/waitgroup/README.md) | [Mutex counter](counter/mutex.md) | [Atomic counter](counter/atomic.md)

```go
package concurrency

// NonStoppingGoRoutine is that a good idea?
func NonStoppingGoRoutine() int {
	atomicCounter := counter.NewAtomicCounter()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			inlinePrint(atomicCounter.Inc())
		}
	}()

	wg.Wait()
	return atomicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutine$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="NonStoppingGoRoutine$" -race 
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
    <td rowspan="3"><img height="320" src="https://media.giphy.com/media/lTrbUqQJCif7NfbXoo/giphy.gif" width="568" alt="?"/></td>
  </tr> 
  <tr>
    <td>No race conditions?</td>
    <td><img height="40" src="../images/yes.png" width="40" alt="?"/></td> 
  </tr>
  <tr>
    <td>Error handling and gracefully shutdown?</td>
    <td><img height="40" src="../images/no.png" width="40" alt="?"/></td>
  </tr>
</tbody>
</table> 

[Next example](example_6.md)
