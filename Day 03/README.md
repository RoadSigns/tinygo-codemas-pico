# Day 03 - Bashing Buttons!

This project will be setting up the board to interact with some LED lights using physical inputs via programmable buttons.

For this exercise you will require the following:

- 1x Raspberry Pi Pico H
- 2x Mini breadboard
- 1x USB to Micro USB Lead
- 1x 5mm Red LED
- 1x 5mm Amber LED
- 1x 5mm Green LED
- 3x 330 ohm resistors
- 11x Male to male jumper wires
- 3x Tall Tactile Buttons
- 3x Button caps

We will roughly be following the official guide, however this guide is in MicroPython. [Day 3 Maker Advent Calendar](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-3-bashing-buttons)

For information on setting up the board and all of the wires, I recommend following the official guide, as this will provide you with the GPIO mapping.

Once everything has been hooked up following the Pihut guide then we can start of writing the activites in TinyGo

## Activites
These activites will be introducing how to use LED lights on the Pico board

### Activity 1: Button Test Program
This actitivy will be just getting a button working with the application and turning a light on once a button has been hit.

Once you have all of the components set up following the guide then we can now look at accessing the button via the GPIO pins.

The first button we will be using is the button connected to GPIO 13.

_(Any button can be started with, I just went with GPIO 13 due to the way I had the compenets lined up infront of me)_

Unlike the LED light, this will be inputting information into the board, so we will be configuring it as an input instead of an output.

```go
inputConfig := machine.PinConfig{
	Mode: machine.PinInputPulldown,
}

buttonOne := machine.GPIO13
buttonOne.Configure(inputConfig)
```

For the mode we have set it to use `machine.PinInputPullDown` instead of `machine.PinInput` this is due to the fact that `PinInput` allow for floating results where the component might not be connected. This will put us in a state where we could be providing invalid states from the button to the application.

Additional information for the `machine.PinChanges` can found here:

- [GPIO Inputs Explained](https://tinygo.org/docs/concepts/peripherals/gpio/#reading-gpio-input)
- [machine package](https://tinygo.org/docs/reference/machine/#gpio)

Now that we have access to the button state, we can now use this to find out if the button is currently in a pushed down state or not.
We can achieve this by accessing the `Get()` method to get the current state of the button. This will return a `true` if the button is currently pushed and `false` if the button isn't being pushed in.

We will be accessing the Green LED light using the same logic as provided in Day 2 acitivites since they are in the same GPIO pins.

From here we can introduce some logical statements such as an `if` statement to create two paths for how what we want to do depending on the state of the button.

```go
if buttonOne.Get() {
	greenLed.High()
	time.Sleep(time.Second)
}
greenLed.Low()
```

If you are unfamiliar with the `if` keyword, additional information can be found here: [Go by example - if statements](https://gobyexample.com/if-else)

Now that we are able to control the light using logical statements, we can add this to a loop to allow us to turn the light on and off depending on the state of the button.

```go
for {
	if buttonOne.Get() {
		greenLed.High()
		time.Sleep(time.Second)
	}
	greenLed.Low()
}
```

If we put this all together then the following script will be:

```go
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
```

### Activity 2: Multiple Button Inputs
Now that we have a working button to LED light example, we can set up the other buttons to the lights, to allow us to control mulitple leds lights with multiple buttons.

For this we will need to set up each of the buttons and get their input. Using the set up document from Pihut, we have our buttons configured to the GPIO of 3, 8 and 13.

```go
inputConfig := machine.PinConfig{
	Mode: machine.PinInputPulldown,
}

buttonOne := machine.GPIO13
buttonOne.Configure(inputConfig)

buttonTwo := machine.GPIO8
buttonOne.Configure(inputConfig)

buttonThree := machine.GPIO3
buttonOne.Configure(inputConfig)
```

Since we have the state of all the buttons now, we can set up the LED lights again using the same logic as we did in Day 2.

```go
outputConfig := machine.PinConfig{
	Mode: machine.PinOutput,
}

greenLed := machine.GPIO18
greenLed.Configure(outputConfig)

yellowLed := machine.GPIO19
yellowLed.Configure(outputConfig)

redLed := machine.GPIO20
redLed.Configure(outputConfig)
```

From this we can use the logical keyword `if` to check if a specific button has been pushed. This will allow us to light up leds depending on the states of individual buttons.

```go
if buttonOne.Get() {
	greenLed.High()
}

if buttonTwo.Get() {
	yellowLed.High()
}

if buttonThree.Get() {
	redLed.High()
}
```

With all of this we can put it into a loop to allow us to manage the state of the LED lights and turn them on and off.

If we put this all together then the complete script:

```go
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

	buttonThree := machine.GPIO3
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
		if buttonOne.Get() {
			greenLed.High()
		}

		if buttonTwo.Get() {
			yellowLed.High()
		}

		if buttonThree.Get() {
			redLed.High()
		}

		time.Sleep(time.Second)

		greenLed.Low()
		yellowLed.Low()
		redLed.Low()
	}
}
```

### Activity 3 - Multiple Button Inputs with else if and else
Now that we are able to access each button and control the LED lights, we can now introduce more complex logical statements like control lights with mulitple buttons.

This is where we will be introducing the `else if` and the `else` keyword.
Using these keywords we will be able to set a default state and then have different combonations of buttons being pressed.

```go
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
```

For the first logical condtion we are using the `&&` which is called an `and` operator. This allows us to only pass `true` to the condition if both sides of the `&&` are both `true`. This means that both `buttonOne` and `buttonTwo` have to be pushed in for the `greenLed` to be set to `High()`.

If no buttons are pushed then we default to the `else` keyword which will just set the `redLed` to `High()`.

We can introduce this into a loop to allow us to update the led lights once they have been pushed.

```go
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
```

### Activity 4: This button or that button?
In the previous activity we learnt about the `&&` operator, in this activity we will be learning about `OR` or `||` operators. This will allow us to run logical statements that allow two buttons to control the same led light.

Unlike the `&&` the `||` operator only requires one side of the logical statement to be `true` for the condition to be `true`.

```go
if buttonOne.Get() || buttonTwo.Get() {
	greenLed.High()
}
greenLed.Low()
```

From here we can now control the `greedLed` with both the `buttonOne` and `buttonTwo`

### Activity 5: Toggling with buttons
Another feature that we will be introducing is the ability to toggle the state of the led from the state of the button.

This can be done by using the `Get()` and `Set()` methods on the `machine.Pin` struct.

```go
greenLed.Set(buttonOne.Get())
```

From here we are now able to turn the light on when we press the button.