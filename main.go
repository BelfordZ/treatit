package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/tty.usbmodem1421")
	servo := gpio.NewServoDriver(firmataAdaptor, "servo", "3")

	work := func() {
		side := true
		gobot.Every(3*time.Second, func() {
			if side == true {
				servo.Max()
			} else {
				servo.Min()
			}
			side = !side
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{servo},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
