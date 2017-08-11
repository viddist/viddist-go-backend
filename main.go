package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/anacrolix/torrent"
)

func main() {
	link := os.Args[1]
	fmt.Println(link)
	var clientConfig torrent.Config
	clientConfig.DataDir = "bt-videos"
	client, err := torrent.NewClient(&clientConfig)
	checkErr(err)
	defer client.Close()

	bootStrapUrls := []string{
		"router.utorrent.com",
		"router.bittorrent.com",
		"dht.transmissionbt.com",
		"dht.aelitis.com",
	}
	//client.AddDHTNodes(bootStraps)
	bootStrapIps := []string{}
	for _, url := range bootStrapUrls {
		ip, err := net.LookupIP(url)
		checkErr(err)
		bootStrapIps = append(bootStrapIps, ip[0].String())
		fmt.Println("IP:", ip[0])
	}
	client.AddDHTNodes(bootStrapIps)
	fmt.Println("Bootstrapped")

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
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}
}
