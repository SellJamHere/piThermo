package models

import (
	"time"

	"github.com/SellJamHere/piBot/thermo"
)

type DbTemperature struct {
	Celsius      *float64   `gorethink:"celsius,omitempty"`
	Fahrenheit   *float64   `gorethink:"fahrenheit,omitempty"`
	MeasuredTime *time.Time `gorethink:"measuredTime,omitempty"`
	DeviceId     *string    `gorethink:"deviceId,omitempty"`
}

func DbTemperatureFromThermo(temp thermo.Temperature) DbTemperature {
	return DbTemperature{
		Celsius:      &temp.Celsius,
		Fahrenheit:   &temp.Fahrenheit,
		MeasuredTime: &temp.MeasuredTime,
	}
}
