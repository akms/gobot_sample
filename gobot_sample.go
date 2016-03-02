package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"time"
)

func main() {

	rgbot := gobot.NewGobot()
	//rgbot.AutoStop = false
	r := raspi.NewRaspiAdaptor("raspi")
	rled := gpio.NewLedDriver(r, "rled", "11")
	bled := gpio.NewLedDriver(r, "bled", "15")
	revent := gobot.NewEvent()

	/*rwork := func() {
		gobot.Every(1*time.Second, func() {
			rled.Toggle()
		})
	}*/

	rwork := func() {
		gobot.On(revent, func(s interface{}) {
			if sint, ok := s.(int); ok {
				fmt.Println("rledstart")
				rled.On()
				fmt.Println("sleep")
				time.Sleep(time.Duration(int64(sint)) * time.Second)
				rled.Off()
				fmt.Println("rledstop")
				fmt.Println("bledstart")
				bled.On()
				fmt.Println("sleep")
				time.Sleep(time.Duration(int64(sint)) * time.Second)
				bled.Off()
				fmt.Println("bledstop")
			}
		})
		gobot.Publish(revent, 5)
		time.Sleep(20 * time.Second)
		gobot.Publish(revent, 1)
		time.Sleep(10 * time.Second)
		gobot.Publish(revent, 5)
	}
	
	/*
	rwork := func() {
		gobot.Once(revent, func(s interface{}) {
			if sint, ok := s.(int); ok {
				fmt.Println("rledstart")
				rled.On()
				fmt.Println("sleep")
				time.Sleep(time.Duration(int64(sint)) * time.Second)
				rled.Off()
				fmt.Println("rledstop")
				fmt.Println("bledstart")
				bled.On()
				fmt.Println("sleep")
				time.Sleep(time.Duration(int64(sint)) * time.Second)
				bled.Off()
				fmt.Println("bledstop")
			}
		})
		gobot.Publish(revent, 5)
	}*/

	/*	bgbot := gobot.NewGobot()
		b := raspi.NewRaspiAdaptor("raspi")
		bled := gpio.NewLedDriver(r, "led", "15")
		bwork := func() {
			gobot.Every(1*time.Second, func() {
				bled.Toggle()
			})
		}*/

	rrobot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{rled, bled},
		rwork,
	)

	/*	brobot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{b},
			[]gobot.Device{bled},
			bwork,
		)
		bgbot.AddRobot(brobot)
		bgbot.Start()
		bgbot.Stop()
	*/
	rgbot.AddRobot(rrobot)
	rgbot.Start()
}
