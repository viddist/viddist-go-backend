package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"log"
	"os"

	"github.com/anacrolix/torrent"
)

func main() {
	link := os.Args[1]

	// Redefining the config is sort of broken atm.
	//var clientConfig torrent.Config
	//clientConfig.DataDir = "bt-videos"
	//client, err := torrent.NewClient(&clientConfig)

	client, err := torrent.NewClient(nil)
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
