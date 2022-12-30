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

	yellowLed := machine.GPIO19
	yellowLed.Configure(outputConfig)

	redLed := machine.GPIO20
	redLed.Configure(outputConfig)

	for {
		if buttonOne.Get() && buttonTwo.Get() {
			greenLed.High()
			redLed.Low()
		} else if buttonOne.Get() {
			yellowLed.High()
			redLed.Low()
		} else {
			redLed.High()
			greenLed.Low()
			yellowLed.Low()
		}

		time.Sleep(time.Second)
	}
}
