# Day 02 - Let’s Get Blinky!

This project will be setting up the board to interact with some LED lights.

For this exercise you will require the following:

- Raspberry Pi Pico H
- Mini breadboard
- USB to Micro USB Lead
- 1x 5mm Red LED
- 1x 5mm Amber LED
- 1x 5mm Green LED
- 3x 330 ohm resistors
- 4x Male to male jumper wires

We will roughly be following the official guide, however this guide is in MicroPython. [Day 2 Maker Advent Calendar](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-2-let-s-get-blinky)

For information on setting up the board and all of the wires, I recommend following the official guide, as this will provide you with the GPIO mapping.

Once everything has been hooked up following the Pihut guide then we can start of writing the activites in TinyGo

## Activites
These activites will be introducing how to use LED lights on the Pico board

### Activity 1: Light each LED
This activity will be turning on the LED lights and then turning them off after a 5 second break. For this we will be introducing the `time` package.

We will be accessing the LED lights via their GPIO pins. This can be done via the `machine` package using their constants. 

Our first LED will be located in Pin 18 so we can access this via `machine.GPIO18`. One we have this we will be able to configure the pin as an output.

```go
outputConfig := machine.PinConfig{
    Mode: machine.PinOutput,
}

greenLed := machine.GPIO18
greenLed.Configure(outputConfig)
```

Now that we have the LED light assigned as a `machine.Pin` we are able to `greenLed.High()` and `greenLed.Low()` to turn the LED light on and off.

This is where we can now control the LED lights, we can introduce the `time` package to turn the lights on and off.

The `time` package allows us to pause the application for a set amount of time.
[time package](https://pkg.go.dev/time)

This can be imported in via the import keyword in the application file.

```go
import (
	"machine"
	"time"
)
```

Now that we have the package imported we are able to pause the time for 5 seconds with the lights on with `time.Sleep(time.Second * 5)`

We can now turn the light on and then back off after five seconds.

```go
greenLed.High()
yellowLed.High()
redLed.High()

time.Sleep(time.Second * 5)

greenLed.Low()
yellowLed.Low()
redLed.Low()
```

If we put this all together then we will be able to configure each LED light, turn them on and then back off after 5 seconds.

```go
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
```

### Activity 2: Flashing LEDs


### Activity 3: LED Sequence