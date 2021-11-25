## Skip List vs Normal List Comparison
### Video (Turkish)
[Skip List Presentation](https://twitter.com/huseyinbabal/status/1463748696580018177)
### How to Run
`go test -bench=.`

### Benchmark Results with Binary Search
```
goos: darwin
goarch: arm64
pkg: github.com/huseyinbabal/skiplist
BenchmarkList-8         	    7393	    159770 ns/op
BenchmarkBinaryList-8   	 8325457	       144.0 ns/op
BenchmarkSkipList-8     	 1295238	       908.0 ns/op
PASS
ok  	github.com/huseyinbabal/skiplist	6.065s
```