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
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Download a magnet link or ipfs hash",
			Action: func(c *cli.Context) error {
				fmt.Println("Downloading", c.Args().First())
				dlMagnet(c.Args().First())
				return nil
			},
		},
	}

	fmt.Println(C.random())

	app.Run(os.Args)
}

func dlMagnet(magnet string) {
	fmt.Println("in the func")
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

	torr, err := client.AddMagnet(magnet)
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
