package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()

	board := edison.NewAdaptor()
	red := gpio.NewLedDriver(board, "3")
	green := gpio.NewLedDriver(board, "5")
	blue := gpio.NewLedDriver(board, "6")

	button := gpio.NewButtonDriver(board, "7")

	enabled := true
	work := func() {
		red.Brightness(0xff)
		green.Brightness(0x00)
		blue.Brightness(0x00)

		flash := false
		on := true

		gobot.Every(50*time.Millisecond, func() {
			if enabled {
				if flash {
					if on {
						red.Brightness(0x00)
						green.Brightness(0xff)
						blue.Brightness(0x00)
						on = false
					} else {
						red.Brightness(0x00)
						green.Brightness(0x00)
						blue.Brightness(0xff)
						on = true
					}
				}
			}
		})

		button.On(gpio.ButtonPush, func(data interface{}) {
			flash = true
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			flash = false
			red.Brightness(0x00)
			green.Brightness(0x00)
			blue.Brightness(0xff)
		})
	}

	robot := gobot.NewRobot(
		"square of fire",
		[]gobot.Connection{board},
		[]gobot.Device{red, green, blue, button},
		work,
	)

	robot.AddCommand("enable", func(params map[string]interface{}) interface{} {
		enabled = !enabled
		return enabled
	})

	master.AddRobot(robot)

	master.Start()
}
