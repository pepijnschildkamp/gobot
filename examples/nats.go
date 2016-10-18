package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/nats"
)

func main() {
	natsAdaptor := nats.NewAdaptorWithAuth("localhost:4222", 1234, "user", "pass")

	work := func() {
		natsAdaptor.On("hello", func(data []byte) {
			fmt.Println("hello")
		})
		natsAdaptor.On("hola", func(data []byte) {
			fmt.Println("hola")
		})
		data := []byte("o")
		gobot.Every(1*time.Second, func() {
			natsAdaptor.Publish("hello", data)
		})
		gobot.Every(5*time.Second, func() {
			natsAdaptor.Publish("hola", data)
		})
	}

	robot := gobot.NewRobot("natsBot",
		[]gobot.Connection{natsAdaptor},
		work,
	)

	robot.Start()
}
