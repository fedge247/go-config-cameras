package main

import (
	"encoding/json"
	"github.com/fedge247/go-config-cameras/hkvision/client"
	"github.com/fedge247/go-config-cameras/hkvision/types"

	"log"
)

/**
 * @author : Donald Trieu
 * @created : 9/25/21, Saturday
**/

func main() {
	newClient, _ := client.NewClient(types.ConstructClient{
		Username: "admin",
		Password: "AI_team123",
		Host: "192.168.1.150",
		Proto: "http",
	})
	//channels , err := newClient.GetStreamChannels()
	//if err != nil {
	//	log.Print(err)
	//}
	//bytes, _ := json.Marshal(channels)
	//log.Printf("Result : %s\n", string(bytes))
	channel, err := newClient.GetSingleStreamChannel("101")
	if err != nil {
		log.Print(err)
	}
	bytes, _ := json.Marshal(channel)
	log.Printf("Result : %s\n", string(bytes))
	//channel.Video.ConstantBitRate = 1024
	//log.Printf("maxFrameRate: %v\n", channel.Video.ConstantBitRate)
	//err = newClient.UpdateSingleStreamChannel(strconv.Itoa(channel.ID), channel)
	//if err != nil {
	//	log.Print(err)
	//}
}