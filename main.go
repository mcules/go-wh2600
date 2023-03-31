package main

import (
	"WH2600/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var mqtt models.Mqtt
	mqttTopic := "YOURE_TOPIC"

	mqtt.Connect("tcp://broker.emqx.io:1883")

	r := gin.Default()
	r.GET("/weatherstation/updateweatherstation.php", func(c *gin.Context) {
		var weatherData models.WeatherData

		c.Bind(&weatherData)
		weatherData.Recalc()

		res, err := json.Marshal(weatherData)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "done",
		})

		mqtt.Publish(mqttTopic, 0, true, string(res))
	})
	r.Run(":8030")

	mqtt.Disconnect()
}
