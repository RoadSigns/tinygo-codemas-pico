# Day 01 - Setting up the board

This project will be setting up the board to use TinyGo and some basic logical
functions in the language to get an understanding of how the language works with
Pico.

We will be roughly following the official guide, however this is in MicroPython.
[Day 1 Marker Advent Calendar](https://thepihut.com/blogs/raspberry-pi-tutorials/maker-advent-calendar-day-1-getting-started)

## Pico H
For these projects we will be using the Pico H, this is just standard Raspberry Pico with the headers preinstalled.

These header pins are called GPIO (General Purpose Input Output) pins.
This allows us to connect all sorts of components to the board to make projects.

We will be using TinyGo to control these and get them to achieve small projects.

A map of the Pico pins can be located here: [Pico GIPO Pins](https://cdn.shopify.com/s/files/1/0176/3274/files/Pico-R3-A4-Pinout_f22e6644-b3e4-4997-a192-961c55fc8cae.pdf?v=1664490511)

## TinyGo
Unlike the official Pihut guide, we will be using TinyGo instead of MicroPython to achieve these projects.

TinyGo is a new compiler for an existing programming language, 
the Go programming language. TinyGo focuses on compiling code written in Go, 
but for smaller kinds of systems such as the Raspberry Pico H.

Much like other languages TinyGo comes with a playground that allow us to work on code in the browser,
so we can get familiar with language before creating projects in it. [TinyGo Playground](https://play.tinygo.org/)

Unlike other languages TinyGo doesn't need to be installed onto the Pico H, we can create the binary on our personal computers
and transfer that across via the `flash` keyword.
[Flashing Binary onto a Pico](https://tinygo.org/docs/reference/microcontrollers/pico/)

## Text editor / IDE
Unlike the official guide, there isn't any official tools like `Thonny` that will allow us to debug, instead we will be
using a mixture of the Playground and a choice of your own text editors/IDE.

## Activities
### Activity 1: Your First Program - Print

For the first day we will just be getting familiar with TinyGo and flashing some information onto the Pico.

The program we will be making will be making the Pico say `This is my Pico talking`.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is my Pico talking")
}
```

For us to achieve this we will need to import the `fmt` package that gives us access to the `PrintLn()` method.

More information on what is available within the `fmt` package can be found here:
- [fmt package in Go](https://pkg.go.dev/fmt)
- [fmt package availability in TinyGo](https://tinygo.org/docs/reference/lang-support/stdlib/#fmt)

Now that we have the code to get Pico H to talk, we just need to flash the device with code.

### Activity 2: Light the Onboard LED

Now that we know how to write code and flash it onto the board. We will be making the onboard LED turn on.

This activity will require us to use the GPIO pins and the breadboard. 

For this we will using the `machine` package to get access to the LED light.
[machine package for pico](https://tinygo.org/docs/reference/microcontrollers/machine/pico/)

The machine package is a package provided by TinyGo that allow us to access parts of the machine without the requirement of having to access each indivdual pin.

From here we will be able to get access to the leds via the `LED` constant, we can then use that to configure the pins.

```go
package main

import (
	"machine"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led.High()
}
```

Once we have flashed this onto the board, the LED will now be shining.

