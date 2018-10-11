```
$ go run test_publisher.go
Run ./test_subscriber -d /ip4/127.0.0.1/tcp/8081/p2p/QmYZJDgeth8jkGWdoFECrLbZuyF5yWEAvyXjMEmvcxRucd
```

In separate console:

```
$ go build test_subscriber.go ## needs to build (not run) for go flag parsing to work
$ ./test_subscriber -d /ip4/127.0.0.1/tcp/8081/p2p/QmYZJDgeth8jkGWdoFECrLbZuyF5yWEAvyXjMEmvcxRucd
&{0xc4200c5040 0xc4202fa2a0 0xc4200a2e60 <nil> 0x141cea0 0x1a1b880 0x1a5f698 60000000000 0xc420082cc0}
```