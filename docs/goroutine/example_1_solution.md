# 1. Let's start with a basic example

[Basic counter](counter/basic.md) | [Race detection](race/race.md)

```go 
package concurrency

// UnexpectedResult what did you expect?
func UnexpectedResult() int {
	basicCounter := counter.NewBasicCounter()

	go func() {
		for i := 0; i < 1000; i++ {
			basicCounter.Inc() //counter++
		}
	}()

	return basicCounter.Count()
}
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="UnexpectedResult$" 
```

```bash
  go test ../../internal/goroutine -v -count=1 -run="UnexpectedResult$" -race 
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
    <td rowspan="3"><img height="360" src="https://media.giphy.com/media/l9TnPq7MBol6E9nn3N/giphy.gif" width="360" alt="?"/></td>
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

[Next example](example_2.md)