# 2. Let's try and fix the issues

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/buildingblocks/sync/waitgroup/README.md)

```go
package concurrency

// UnexpectedResultFix is it fixed?
func UnexpectedResultFix() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			basicCounter.Inc()
		}
	}()

	wg.Wait()
	return basicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="UnexpectedResultFix$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="UnexpectedResultFix$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/Od0QRnzwRBYmDU3eEO/giphy.gif" width="360" alt="?"/></td>
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

[Next example](example_3.md)