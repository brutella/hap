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

	a := accessory.NewDimmerColorTemperatureLightbulb(accessory.Info{
		Name: "DimmerColorTemperatureLightbulb",
	})

	a.Lightbulb.Brightness.SetMaxValue(100)
	a.Lightbulb.Brightness.SetMinValue(0)
	a.Lightbulb.Brightness.SetStepValue(1)
	a.Lightbulb.ColorTemperature.SetMaxValue(300)
	a.Lightbulb.ColorTemperature.SetMinValue(200)
	a.Lightbulb.ColorTemperature.SetStepValue(1)

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

	a.Lightbulb.Brightness.OnValueRemoteUpdate(func(v int) {
		log.Println("DimmerColorTemperatureLightbulb change brightness,value is " + strconv.Itoa(v))
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
