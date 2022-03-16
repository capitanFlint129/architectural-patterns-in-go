# Channel multiplexer

OrRecursion accepts an arbitrary number of channels and returns one channel that returns the events any of the internal:

```
func OrRecursion (channels ...<- chan interface{}) <- chan interface{} {
	// ...
}
```

More: https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html
