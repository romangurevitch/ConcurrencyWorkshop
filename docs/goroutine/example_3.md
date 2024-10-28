# 3. Let's make a small change :)

[Basic counter](counter/basic.md) | [Race detection](race/race.md) | [WaitGroup](../../internal/buildingblocks/sync/waitgroup/README.md)

```go
package concurrency

// LetsMakeASmallChange ohh no!
func LetsMakeASmallChange() int {
	basicCounter := counter.NewBasicCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			basicCounter.Inc()
		}()
	}

	wg.Wait()
	return basicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="LetsMakeASmallChange$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="LetsMakeASmallChange$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/jBncVXsbkjB8Nv5Igm/giphy.gif" width="360" alt="?"/></td>
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

[Solution](example_3_solution.md)