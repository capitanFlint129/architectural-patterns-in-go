# Channel multiplexer

The "or" function accepts an arbitrary number of channels and returns one channel that returns the events of any of the internal:

```
func or (channels ...<- chan interface{}) <- chan interface{} {
	// ...
}
```
