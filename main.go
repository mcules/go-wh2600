package main

import (
	"WH2600/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	config "github.com/spf13/viper"
	"net/http"
)

func main() {
	var mqtt models.Mqtt
	var err error

	// get config from env file
	config.SetConfigFile(".env")
	if err = config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// connect to mqtt broker
	mqtt.Connect(config.Get("MQTT_PROTOCOL").(string)+"://"+config.Get("MQTT_BROKER").(string)+":"+config.Get("MQTT_PORT").(string), config.Get("MQTT_CLIENT_ID").(string))

	// create webserver routes
	r := gin.Default()
	r.GET("/weatherstation/updateweatherstation.php", func(c *gin.Context) {
		var weatherData models.WeatherData

		if err := c.Bind(&weatherData); err != nil {
			fmt.Println(err)
		}
		weatherData.Recalc()

		res, err := json.Marshal(weatherData)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

		// publish data to mqtt
		mqtt.Publish(config.Get("MQTT_TOPIC").(string), 0, true, string(res))
	})
	// start webserver
	if err := r.Run(":" + config.Get("WEBSERVER_PORT").(string)); err != nil {
		fmt.Println(err)
	}

	// close mqtt connection
	mqtt.Disconnect()
}
