```
$ go run test_publisher.go
Run ./test_subscriber -d /ip4/127.0.0.1/tcp/8081/p2p/QmYZJDgeth8jkGWdoFECrLbZuyF5yWEAvyXjMEmvcxRucd
```

In separate console:

```
$ go build test_subscriber.go ## needs to build (not run) for go flag parsing to work
$ ./test_subscriber -d /ip4/127.0.0.1/tcp/8081/p2p/QmYZJDgeth8jkGWdoFECrLbZuyF5yWEAvyXjMEmvcxRucd
```
