package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SellJamHere/dataWhereHouse/models"
	"github.com/SellJamHere/piBot/thermo"
	r "github.com/dancannon/gorethink"
)

const (
	piSerial   = ""
	piDeviceId = "e0a701e4-f37c-4d6b-adbb-4cc6995a266b"
)

func main() {
	var session *r.Session

	tempReader, err := thermo.NewTemperatureReader(piSerial)
	if err != nil {
		log.Fatalln(err.Error())
	}

	session, err = r.Connect(r.ConnectOpts{
		Address: "162.243.156.65:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		temp, err := tempReader.ReadTemp()
		if err != nil {
			fmt.Println(err)
			panic("Error reading temp")
		}

		dbTemp := models.DbTemperatureFromThermo(*temp)

		_, err = r.DB("dataWhereHouse").Table("event").Insert(dbTemp).RunWrite(session)
		if err != nil {
			log.Fatalln(err.Error())
		}
		time.Sleep(5 * time.Minute)
	}
}
