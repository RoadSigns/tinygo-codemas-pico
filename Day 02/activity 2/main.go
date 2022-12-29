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

	counter := 0

	for counter <= 11 {
		greenLed.High()
		yellowLed.High()
		redLed.High()
		time.Sleep(time.Second)

		greenLed.Low()
		yellowLed.Low()
		redLed.Low()
		time.Sleep(time.Second)

		counter++
	}
}
