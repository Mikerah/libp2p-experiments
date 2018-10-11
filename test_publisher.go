package main

import (
	"context"
	mrand "math/rand"
	"github.com/libp2p/go-floodsub"

	"github.com/libp2p/go-libp2p"
	"fmt"
	"time"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sourcePort := 8081
	fmt.Println("starting server on ", sourcePort)
	rand := mrand.New(mrand.NewSource(int64(sourcePort)))
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand)
	if err != nil {
		panic(err)
	}
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", sourcePort))

	h, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		panic(err)
	}

	var port string
	for _, la := range h.Network().ListenAddresses() {
		fmt.Println("in for loop")
		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			port = p
			break
		}
	}

	if port == "" {
		panic("was not able to find actual local port")
	}


	fmt.Printf("Run ./test_subscriber -d /ip4/127.0.0.1/tcp/%v/p2p/%s\n", port, h.ID().Pretty())

	pubsub, err := floodsub.NewFloodSub(ctx, h)
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

	<-quit

	//ch, err := pubsub.Subscribe("car")
	//if err != nil {
	//	panic(err)
	//}
	//
	//for {
	//	newmsg, err := ch.Next(ctx)
	//	if(err != nil ) {
	//		panic(err)
	//	}
	//	fmt.Println("message is ", newmsg)
	//}

	fmt.Println("done")
}
