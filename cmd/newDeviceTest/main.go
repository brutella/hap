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

	// a := accessory.NewColorTemperatureLightbulb(accessory.Info{
	// 	Name: "ColorTemperatureLightbulb",
	// })

	// b := accessory.NewThermostat(accessory.Info{
	// 	Name: "Thermostat",
	// })

	e := accessory.NewHeater(accessory.Info{
		Name: "Heater",
	})

	e.Heater.CurrentHeaterCoolerState.OnSetRemoteValue(func(v int) error {
		log.Println("CurrentHeaterCoolerState:" + strconv.Itoa(v))
		return nil
	})

	e.Heater.CurrentTemperature.OnSetRemoteValue(func(v float64) error {
		log.Println("CurrentTemperature:" + strconv.Itoa(int(v)))
		return nil
	})

	e.Heater.HeatingThresholdTemperature.OnSetRemoteValue(func(v float64) error {
		log.Println("HeatingThresholdTemperature:" + strconv.Itoa(int(v)))
		return nil
	})

	e.Heater.TargetHeaterCoolerState.OnSetRemoteValue(func(v int) error {
		log.Println("TargetHeaterCoolerState:" + strconv.Itoa(v))
		return nil
	})

	e.Heater.Active.OnSetRemoteValue(func(v int) error {
		log.Println("Active:" + strconv.Itoa(v))
		return nil
	})

	// f := accessory.NewAirPurifier(accessory.Info{
	// 	Name: "AirPurifier",
	// })

	// h := accessory.NewWindowCovering(accessory.Info{
	// 	Name: "WindowCovering",
	// })

	// a1 := accessory.NewTemperatureSensor(accessory.Info{
	// 	Name: "TemperatureSensor",
	// })

	// a2 := accessory.NewHumidifier(accessory.Info{
	// 	Name: "Humidifier",
	// })

	// a1.TempSensor.CurrentTemperature.SetValue(30.5)

	// a3 := accessory.NewHumiditySensor(accessory.Info{
	// 	Name: "HumiditySensor",
	// })

	// a3.HumiditySensor.CurrentRelativeHumidity.SetValue(27.5)

	// a4 := accessory.NewLightSensor(accessory.Info{
	// 	Name: "Light Sensor",
	// })

	// a4.LightSensor.CurrentAmbientLightLevel.SetValue(200)

	// a5 := accessory.NewAirQualitySensor(accessory.Info{
	// 	Name: "AirQualitySensor",
	// })
	// a5.AirQualitySensor.AirQuality.SetValue(2)

	// a6 := accessory.NewCarbonDioxideSensor(accessory.Info{
	// 	Name: "CarbonDioxideSensor",
	// })

	// a6.CarbonDioxideSensor.CarbonDioxideDetected.SetValue(1)
	// a6.CarbonDioxideSensor.CarbonDioxideLevel.SetValue(300)

	// a7 := accessory.NewMotionSensor(accessory.Info{
	// 	Name: "MotionSensor",
	// })

	// a7.MotionSensor.MotionDetected.SetValue(true)

	// a8 := accessory.NewContactSensor(accessory.Info{
	// 	Name: "ContactSensor",
	// })

	// a8.ContactSensor.ContactSensorState.SetValue(1)

	// a9 := accessory.NewOccupancySensor(accessory.Info{
	// 	Name: "OccupancySensor",
	// })

	// a9.OccupancySensor.OccupancyDetected.SetValue(1)

	a10 := accessory.NewHeater_New(accessory.Info{
		Name: "Heater New",
	})

	a10.Heater.Active.OnSetRemoteValue(func(v int) error {
		log.Println("Active:" + strconv.Itoa(v))
		return nil
	})

	a10.Heater.CurrentHeaterCoolerState.OnSetRemoteValue(func(v int) error {
		log.Println("CurrentHeaterCoolerState:" + strconv.Itoa(v))
		return nil
	})

	a10.Heater.CurrentTemperature.OnSetRemoteValue(func(v float64) error {
		log.Println("CurrentTemperature:" + strconv.Itoa(int(v)))
		return nil
	})

	a10.Heater.HeatingThresholdTemperature.OnSetRemoteValue(func(v float64) error {
		log.Println("HeatingThresholdTemperature:" + strconv.Itoa(int(v)))
		return nil
	})

	a10.Heater.CurrentTemperature.SetValue(10)

	//s, err := hap.NewServer(hap.NewFsStore("./db"), d.A, a.A, b.A, e.A, f.A, h.A, a1.A, a2.A, a3.A, a4.A, a5.A, a6.A, a7.A, a8.A, a9.A)
	s, err := hap.NewServer(hap.NewFsStore("./db"), d.A, e.A, a10.A)
	if err != nil {
		log.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	// a.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
	// 	if on == true {
	// 		log.Println("Client changed switch to on")
	// 	} else {
	// 		log.Println("Client changed switch to off")
	// 	}
	// })

	// a.Lightbulb.ColorTemperature.OnValueRemoteUpdate((func(v int) {
	// 	log.Println("DimmerColorTemperatureLightbulb change ColorTemperature,value is " + string(v))
	// }))

	// b.Thermostat.CurrentHeatingCoolingState.OnSetRemoteValue(func(v int) error {
	// 	log.Println("Thermostat change CurrentHeatingCoolingState,value is " + string(v))
	// 	return nil
	// })

	// b.Thermostat.TargetHeatingCoolingState.OnSetRemoteValue(func(v int) error {
	// 	log.Println("Thermostat change TargetHeatingCoolingState,value is " + string(v))
	// 	return nil
	// })

	// b.Thermostat.CurrentTemperature.OnSetRemoteValue(func(v float64) error {
	// 	log.Println("Thermostat change CurrentTemperature,value is " + string(int(v)))
	// 	return nil
	// })

	// b.Thermostat.TargetTemperature.OnSetRemoteValue(func(v float64) error {
	// 	log.Println("Thermostat change TargetTemperature,value is " + string(int(v)))
	// 	return nil
	// })

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
