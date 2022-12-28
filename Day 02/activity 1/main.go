package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{
		Mode: machine.PinOutput,
	}

	greenLed := machine.GPIO18
	greenLed.Configure(outputConfig)

	yellowLed := machine.GPIO19
	yellowLed.Configure(outputConfig)

	redLed := machine.GPIO20
	redLed.Configure(outputConfig)

	greenLed.High()
	yellowLed.High()
	redLed.High()

	time.Sleep(time.Second * 5)

	greenLed.Low()
	yellowLed.Low()
	redLed.Low()

}