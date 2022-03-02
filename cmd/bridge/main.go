package main

import (
	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/log"

	"context"
	syslog "log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a := accessory.NewOutlet(accessory.Info{
		Name:         "My Test Device",
		SerialNumber: "1234",
		Model:        "a",
		Manufacturer: "Matthias",
		Firmware:     "1.0",
	})

	l := accessory.NewLightbulb(accessory.Info{
		Name:         "My Light bulb",
		SerialNumber: "1234",
		Model:        "a",
		Manufacturer: "Matthias",
		Firmware:     "1.0",
	})

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A, l.A)
	if err != nil {
		log.Info.Panic(err)
	}

	mylogger := syslog.New(os.Stdout, "BLUB ", syslog.LstdFlags|syslog.Lshortfile)
	log.Debug = &log.Logger{mylogger}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		signal.Stop(c) // stop delivering signals
		cancel()
	}()

	s.ListenAndServe(ctx)
}
