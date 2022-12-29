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
Now that we can turn the LED lights on and off, we will be looking at introducing a loop that will allow us to continuously flash the LED lights.

In the example they use a `while` loop in Python, however in Go they have different loops.

We will be using a `for` loop with a condition to achieve the same result as the `while` loop. This will keep looping the as long as the condition is `true`.

```go
for condition {

}
```

To achieve this for loop we will need to assign a value to a variable so we can keep the state of the counter for the loop.

This is achieved in Go by using `:=`
```go
counter := 0
```

This allows us to increment the counter everytime we loop over the lights flashing.

Now that we have the loop, the LED logic and the counter variable, we can start looping and counting up the counter until the logic is false.

```go
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
```

At the end of the loop we are using incrementing the counter variable by 1 by using `++`. This ensures that we aren't stuck in an infinite loop since the counter will increment to 11.


The final code looks like the following:
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
```

### Activity 3: LED Sequence
For the final activity this day, we will be creating a traffic light system using the 3 LED lights.

This will be using leaning on the `time` package, `for` loops and introducing `structs` which is a type in Go.

First we will be creating the `TrafficLight` struct that will contain all of the commands to put the lights into the state that we want.

```go
type TrafficLights struct {
	red   machine.Pin
	amber machine.Pin
	green machine.Pin
}
```

Each property on the traffic lights will be a `machine.Pin`, this is due to us providing direct access to the LED lights to the traffic light.

That we have a struct with access to the pins, we will be making methods on the struct to control these. Methods are like functions but allow us to encasplate the logic so it can only be access from the public APIs on the struct that we provide.

For the `TrafficLight` struct we will be making the methods `stop`, `start`, `slow` and `warn` publicly available to control the state of the lights.

We have made all of the lights private inside of the `TrafficLights` struct, so we can't get ourselves into a state where we can manually change them and set the lights to red and green being on and amber being off, since this isn't a valid traffic light state.

Information around how we make struct properties public and private can be found here: [Go by Example: Structs](https://gobyexample.com/structs)

Now that we have the struct, we need to create the methods defined.
[Go by Example: Methods](https://gobyexample.com/methods)
```go
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
```

Here we are able to set the exact state of the lights depending on the situation we want to occur. So for example when we want to `stop()` then all the lights apart from `red` is set to `Low()` while `red` is set to `High()`.

This allows us to reuse the method to always set the lights to a stop state.

With the new struct and all of the methods to control the lights in the exact state that we want, we can create a loop that will allow us to loop over the states of a traffic light.

```go
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
```

Using the `time` package we are able to give cars time to move and even stop of traffic coming the other way, if we had another set of lights available.

The full activity is the following:
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
```