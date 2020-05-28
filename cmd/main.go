package main

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ryusei112/temp-sensor-go/i2c"
	"github.com/ryusei112/temp-sensor-go/tls"
)

const (
	ThingName = ""
	Endpoint  = ""
	PubTopic  = "topic/to/publish"
	QoS       = 1
)

func main() {

	log.Println("IoT Endpoint:%s", Endpoint)

	tlsConfig, err := tls.SetCredentials()
	if err != nil {
		panic(fmt.Sprintf("failed to construct TLS Config: :%v", err))
	}

	// Connect
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", Endpoint, 8883))
	opts.SetClientID(ThingName).SetTLSConfig(tlsConfig)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("failed to connect broker: :%v", token.Error()))
	}
	defer client.Disconnect(250)

	data := i2c.GetRoomEnv()
	data.Device = ThingName

	j, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	pubMsg := string(j)
	log.Printf("Created JSON: %s\n", pubMsg)

	// Publish
	log.Printf("publishing %s...\n", PubTopic)
	if token := client.Publish(PubTopic, QoS, false, pubMsg); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("failed to publish %s :%v", PubTopic, token.Error()))
	}
	log.Println("Message published")
}
