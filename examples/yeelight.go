package main

import (
  "fmt"
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
)

func main() {
	gbot := gobot.NewGobot()

	bleAdaptor := ble.NewBLEClientAdaptor("ble", os.Args[1])
	yeelight := ble.NewYeelightDriver(bleAdaptor, "yeelight")

	work := func() {
    yeelight.On()
  	// r := uint8(255)
  	// g := uint8(255)
  	// b := uint8(255)
  	// _ = yeelight.SetRGB(r, g, b)
		gobot.Every(5*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			_ = yeelight.SetRGB(r, g, b)
      colour := fmt.Sprintf("Changing Colour to: %d,%d,%d", r, g, b)
      fmt.Println(colour)
		})
	}

	robot := gobot.NewRobot("yeelightBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{yeelight},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
