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

	buttonTwo := machine.GPIO8
	buttonOne.Configure(inputConfig)

	outputConfig := machine.PinConfig{
		Mode: machine.PinOutput,
	}

	greenLed := machine.GPIO18
	greenLed.Configure(outputConfig)

	for {
		if buttonOne.Get() || buttonTwo.Get() {
			greenLed.High()
		}
		greenLed.Low()

		time.Sleep(time.Second)
	}
}
