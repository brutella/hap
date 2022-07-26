# hap

[![GoDoc Widget]][GoDoc] [![Travis Widget]][Travis]

`hap` (previously [hc](https://github.com/brutella/hc)) is a lightweight library to develop HomeKit accessories in Go.
It abstracts the **H**omeKit **A**ccessory **P**rotocol (HAP) and makes it easy to work with [services](service/README.md) and [characteristics](characteristic).

`hap` handles the underlying communication between HomeKit accessories and clients.
You can focus on implementing the business logic for your accessory, without having to worry about the protocol.

Here are some projects which use `hap`.

- [hkknx](https://hochgatterer.me/hkknx)
- [hkcam](https://github.com/brutella/hkcam)

**What is HomeKit?**

[HomeKit][homekit] is a set of protocols and libraries from Apple. It is used by Apple's platforms to communicate with smart home appliances. A non-commercial version of the documentation is now available on the [HomeKit developer website](https://developer.apple.com/homekit/).

HomeKit is fully integrated into iOS since iOS 8. Developers can use [HomeKit.framework](https://developer.apple.com/documentation/homekit) to communicate with accessories using high-level APIs.

<img alt="Home+.app" src="_img/home+.png?raw=true" width="150" />

I've developed the [Home+][home+] app to control HomeKit accessories from iPhone, iPad, and Apple Watch.
If you want to support `hap`, please purchase Home from the [App Store][home-appstore]. That would be awesome. ❤️

[home+]: https://hochgatterer.me/home+/
[home-appstore]: http://itunes.apple.com/app/id995994352
[GoDoc]: https://godoc.org/github.com/brutella/hap
[GoDoc Widget]: https://godoc.org/github.com/brutella/hap?status.svg
[Travis]: https://travis-ci.org/brutella/hap
[Travis Widget]: https://travis-ci.org/brutella/hap.svg

**Migrate from `hc`**

This library is a rewrite of [hc](https://github.com/brutella/hc).
If you want to migrate from `hc`, consider the following changes.

- Instead of `hc.NewIPTransport(...)` you now call [hap.NewServer(...)](https://pkg.go.dev/github.com/brutella/hap#NewServer) to create a server.
- You can create your own persistent storage by implementing the [Store](store.go) interface.
- Setting the value of a characteristic can now fail. Fixes [hc#163](https://github.com/brutella/hc/issues/163)
- You can define custom http handlers. Fixes [hc#212](https://github.com/brutella/hc/issues/212)
```go
server.ServeMux().HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("pong"))
})
```
- You can define your own public and private key (just in case) by setting the [Key](https://github.com/brutella/hap/blob/master/server.go#L42) field of the server. Otherwise those keys are generate and stored on disk for you.
```go
server.Key = hap.KeyPair{
	Public:  []byte{...},
	Private: []byte{...},
}
```
- The base structs for accessories, services and characteristics are now [accessory.A](accessory/a.go), [service.S](service/s.go), [characteristic.C](characteristic/c.go)

## Features

- Supports Go modules (requires Go 1.13)
- Full implementation of the HAP in Go
- Supports all HomeKit [services](service) and [characteristics](characteristic)
- Built-in service announcement via DNS-SD using [dnssd](http://github.com/brutella/dnssd)
- Runs on linux and macOS
- Documentation: http://godoc.org/github.com/brutella/hap

## Usage

In a following example a simple on/off switch is created.
It can be paired with HomeKit using the Apple Home app – use the pin code *00102003*.

```go
package main

import (
	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create the switch accessory.
	a := accessory.NewSwitch(accessory.Info{
		Name: "Lamp",
	})

	// Store the data in the "./db" directory.
	fs := hap.NewFsStore("./db")

	// Create the hap server.
	server, err := hap.NewServer(fs, a.A)
	if err != nil {
		// stop if an error happens
		log.Panic(err)
	}

	// Setup a listener for interrupts and SIGTERM signals
	// to stop the server.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		// Stop delivering signals.
		signal.Stop(c)
		// Cancel the context to stop the server.
		cancel() 
	}()

	// Run the server.
	server.ListenAndServe(ctx)
}
```

### Events

The library provides callback functions, which let you know when a client updates a characteristic value.
The following example shows how to get notified when the [On](characteristic/on.go) characteristic value changes.

```go
a.Switch.On.OnValueRemoteUpdate(func(on bool) {
    if on == true {
        log.Println("Switch is on")
    } else {
        log.Println("Switch is off")
    }
})
```

If you want to change the state of a switch programmatically, you call [SetValue(...)](https://pkg.go.dev/github.com/brutella/hap/characteristic#Bool.SetValue).

```go
a.Switch.On.SetValue(true)
```

The library takes care of the rest and notifies all connected clients that the state has changed.

## Multiple Accessories

When you create a server you can specify multiple accessories like this.

```go
var a1, a2, a3 *accessory.A
s, err := hap.NewServer(fs, a1, a2, a3)
```

By doing so, the first accessory `a1` appears as a bridge in HomeKit.
When adding the accessories to HomeKit, iOS only shows the bridge accessory.
Once the bridge was added, the other accessories appear automatically.

HomeKit requires that every accessory has a unique id, which must not change between system restarts.
`hap` automatically assigns the ids for you based on the order in which the accessories are added to the server.

The best would be to specify the unique id for every accessory yourself, like this

```go
a1.Id = 1
a2.Id = 2
```

## Accessory Architecture

HomeKit uses a hierarchical architecture to define accessories, services and characeristics.
At the root level there is an accessory.
Every accessory contains services.
And every service contains characteristics.

For example a [lightbulb accessory](accessory/lightbulb.go) contains a [lightbulb service](service/lightbulb.go).
This service contains the [on](characteristic/on.go) characteristic.

There are predefined accessories, services and characteristics available in HomeKit.
Those types are defined in the packages [accessory](accessory), [service](service), [characteristic](characteristic).

# Contact

Matthias Hochgatterer

Website: [https://hochgatterer.me](https://hochgatterer.me)

Github: [https://github.com/brutella](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

`hap` is available under the Apache License 2.0 license. See the LICENSE file for more info.

[homekit]: https://developer.apple.com/homekit/
