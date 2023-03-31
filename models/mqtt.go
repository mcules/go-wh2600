package models

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	Client                paho.Client
	Token                 paho.Token
	MessagePubHandler     paho.MessageHandler
	ConnectHandler        paho.OnConnectHandler
	ConnectionLostHandler paho.ConnectionLostHandler
}

func (mqtt *Mqtt) Connect(broker string) {
	mqtt.MessagePubHandler = func(client paho.Client, msg paho.Message) {
		fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
	}
	mqtt.ConnectHandler = func(client paho.Client) {
		fmt.Println("Connected")
	}
	mqtt.ConnectionLostHandler = func(client paho.Client, err error) {
		fmt.Printf("Connection Lost: %s\n", err.Error())
	}

	options := paho.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(mqtt.MessagePubHandler)
	options.OnConnect = mqtt.ConnectHandler
	options.OnConnectionLost = mqtt.ConnectionLostHandler

	mqtt.Client = paho.NewClient(options)
	mqtt.Token = mqtt.Client.Connect()
	if mqtt.Token.Wait() && mqtt.Token.Error() != nil {
		panic(mqtt.Token.Error())
	}
}

func (mqtt *Mqtt) Publish(topic string, qos byte, retained bool, message string) {
	mqtt.Token = mqtt.Client.Publish(topic, qos, retained, message)
	mqtt.Token.Wait()
}

func (mqtt *Mqtt) Disconnect() {
	mqtt.Client.Disconnect(100)
}
