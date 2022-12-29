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

	trafficLights := TrafficLights{
		red:   redLed,
		amber: yellowLed,
		green: greenLed,
	}

	for {
		trafficLights.stop()
		time.Sleep(time.Second * 5)
		trafficLights.ready()
		time.Sleep(time.Second * 2)
		trafficLights.start()
		time.Sleep(time.Second * 5)
		trafficLights.slow()
		time.Sleep(time.Second * 2)
	}
}

type TrafficLights struct {
	red   machine.Pin
	amber machine.Pin
	green machine.Pin
}

func (t *TrafficLights) stop() {
	t.red.High()
	t.amber.Low()
	t.green.Low()
}

func (t *TrafficLights) start() {
	t.red.Low()
	t.amber.Low()
	t.green.High()
}

func (t *TrafficLights) ready() {
	t.red.Low()
	t.amber.High()
	t.green.Low()
}

func (t *TrafficLights) slow() {
	t.red.Low()
	t.amber.High()
	t.green.High()
}
