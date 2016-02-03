package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/SellJamHere/piBot/thermo"
	"github.com/SellJamHere/piThermo/models"
	r "github.com/dancannon/gorethink"
)

const (
	serverAddress = ""
	piSerial      = ""
	piDeviceId    = "e0a701e4-f37c-4d6b-adbb-4cc6995a266b"
	diffThreshold = 0.15
)

func main() {
	var session *r.Session

	tempReader, err := thermo.NewTemperatureReader(piSerial)
	if err != nil {
		log.Fatalln(err.Error())
	}

	session, err = r.Connect(r.ConnectOpts{
		Address: serverAddress,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	var lastTemp float64
	lastTemp = 0.0

	for {
		temp, err := tempReader.ReadTemp()
		if err != nil {
			fmt.Println(err)
			panic("Error reading temp")
		}

		diff := temp.Celsius - lastTemp

		if math.Abs(diff) > diffThreshold {
			lastTemp = temp.Celsius

			dbTemp := models.DbTemperatureFromThermo(*temp)
			dbTemp.DeviceId = piDeviceId

			_, err = r.DB("dataWhereHouse").Table("event").Insert(dbTemp).RunWrite(session)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
