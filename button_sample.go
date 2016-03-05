package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {

	gbot := gobot.NewGobot()
	r := raspi.NewRaspiAdaptor("raspi")
	button := gpio.NewButtonDriver(r, "button", "11")
	led := gpio.NewLedDriver(r, "led", "15")

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			led.On()
		})
		gobot.On(button.Event("release"), func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
