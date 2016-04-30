# nstat

Go library clone of [nstat](https://github.com/shemminger/iproute2/blob/master/misc/nstat.c)

```go
c := &nstat.Counters{DumpZeros: true}
for counter, value := range c.Get() {
	fmt.Printf("%-32s %d\n", counter, value)
}
```

see [API documentation](http://godoc.org/github.com/sparrc/nstat)

## Installing the binary:

```
go get github.com/sparrc/nstat/...
# -a is always on:
nstat -z
```