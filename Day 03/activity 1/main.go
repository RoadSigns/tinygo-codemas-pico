package main

import (
	"machine"
	"time"
)

func main() {
	inputConfig := machine.PinConfig{
		Mode: machine.PinInputPulldown,
	}

	buttonOne := machine.GPIO13
	buttonOne.Configure(inputConfig)

	outputConfig := machine.PinConfig{
		Mode: machine.PinOutput,
	}

	greenLed := machine.GPIO18
	greenLed.Configure(outputConfig)

	redLed := machine.GPIO20
	redLed.Configure(outputConfig)

	for {
		if buttonOne.Get() {
			greenLed.High()
			time.Sleep(time.Second)
		}
		greenLed.Low()
	}
}
