# Channel multiplexer

The "or" function accepts an arbitrary number of channels and returns one channel that returns the events of any of the internal:

```
func or (channels ...<- chan interface{}) <- chan interface{} {
	// ...
}
```

More: https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html
