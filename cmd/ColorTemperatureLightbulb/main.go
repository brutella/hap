// This example show an example of a switch accessory
// which periodically changes it's state between on and off.
package main

import (
	"strconv"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	d := accessory.NewBridge(accessory.Info{
		Name: "Bridge",
	})

	a := accessory.NewColorTemperatureLightbulb(accessory.Info{
		Name: "ColorTemperatureLightbulb",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), d.A, a.A)
	if err != nil {
		log.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	a.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("Client changed switch to on")
		} else {
			log.Println("Client changed switch to off")
		}
	})

	a.Lightbulb.ColorTemperature.OnValueRemoteUpdate((func(v int) {
		log.Println("DimmerColorTemperatureLightbulb change ColorTemperature,value is " + strconv.Itoa(v))
	}))

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		signal.Stop(c)
		cancel()
	}()

	s.ListenAndServe(ctx)
}
