package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"log"
	"os"

	"github.com/anacrolix/dht"
	"github.com/anacrolix/torrent"
)

func main() {
	link := os.Args[1]

	addrs, err := dht.GlobalBootstrapAddrs()
	checkErr(err)
	var clientConfig torrent.Config
	var dhtConfig dht.ServerConfig
	dhtConfig.StartingNodes = addrs
	clientConfig.DataDir = "bt-videos"
	clientConfig.DHTConfig = dhtConfig
	client, err := torrent.NewClient(&clientConfig)
	// I'm not entirely sure this config is working since eduroam seems
	// to mess with the dht. Test someplace else.
	//client, err := torrent.NewClient(nil)
	checkErr(err)
	defer client.Close()

	torr, err := client.AddMagnet(link)
	checkErr(err)
	<-torr.GotInfo()
	fmt.Println("Torrent info:", torr.Info().Files)
	fmt.Println("Added link")
	torr.DownloadAll()

	if client.WaitAll() {
		log.Print("Downloaded the torrent")
	} else {
		log.Fatal("Did not finish downloading for some reason")
	}

	fmt.Println(C.random())
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}
}
