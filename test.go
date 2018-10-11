package main

import (
	"context"
	"github.com/libp2p/go-floodsub"

	"github.com/libp2p/go-libp2p"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("host %+v", h)
	fmt.Println("tsra")
	fmt.Println(h.ID().Pretty())
	fmt.Printf("connect to /ip4/127.0.0.1/tcp/%v/p2p/%s\n", 3001, h.ID().Pretty())

	pubsub, err := floodsub.NewFloodSub(ctx, h)
	if err != nil {
		panic(err)
	}

	ch, err := pubsub.Subscribe("foobar")
	if err != nil {
		panic(err)
	}

	msg := []byte("hello world")

	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				err = pubsub.Publish("foobar", msg)
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()

	for {
		newmsg, err := ch.Next(ctx)
		if(err != nil ) {
			panic(err)
		}
		fmt.Println("message is ", newmsg)
	}

	fmt.Println("done")
}
