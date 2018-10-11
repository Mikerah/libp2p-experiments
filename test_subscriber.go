package main

import (
	"github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p"
	"context"
	"fmt"
	"github.com/libp2p/go-floodsub"
	"flag"
)

func main() {
	dest := flag.String("d", "", "Destination multiaddr string")
	flag.Parse()
	if *dest == "" {
		panic("Please provide d flag")
	}
	maddr, err := multiaddr.NewMultiaddr(*dest)
	if err != nil {
		panic(err)
	}
	info, err := peerstore.InfoFromP2pAddr(maddr)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	host, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}
	host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)
    fmt.Println(host)

	pubsub, err := floodsub.NewFloodSub(ctx, host)
	if err != nil {
		panic(err)
	}

	ch, err := pubsub.Subscribe("foobar")
	if err != nil {
		panic(err)
	}

	for {
		newmsg, err := ch.Next(ctx)
		if(err != nil ) {
			panic(err)
		}
		fmt.Println("message is ", newmsg)
	}

	fmt.Println("done")
}


