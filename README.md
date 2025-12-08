## hamqtt

[![Go Reference](https://pkg.go.dev/badge/github.com/nathan-osman/hamqtt.svg)](https://pkg.go.dev/github.com/nathan-osman/hamqtt)
[![MIT License](https://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](https://opensource.org/licenses/MIT)

This package aims to provide an easy interface for exposing entities to [Home Assistant](https://www.home-assistant.io/) via [MQTT](https://mqtt.org/).

### Installation

Adding hamqtt to your application is as easy as:

```golang
import "github.com/nathan-osman/hamqtt"
```

### Usage

To expose entities to Home Assistant, you must first create a connection (`Conn`):

```golang
c, err := hamqtt.New(&hamqtt.Config{
    Addr:     "1.2.3.4:1883",
    Username: "myusername",
    Password: "password123",
})
if err != nil {
    panic(err)
}
defer c.Close()
```

From there, you can create entities by calling their respective member function. For example, to create a light:

```golang
l, err := c.Light(
    &hamqtt.EntityConfig{
        ID:   "mylight",
        Name: "My Light",
    },
    &hamqtt.LightConfig{
        ChangeCallback: func(on bool) bool {
            if on {
                fmt.Println("Light turned on")
            } else {
                fmt.Println("Light turned off")
            }
            return true
        },
    },
)
```

This will create a corresponding entity for your light that you can then control directly from the Home Assistant UI:

<img src="https://github.com/nathan-osman/hamqtt/blob/main/dist/example-light.png?raw=true" width="229" alt="Screenshot of light control in Home Assistant" />

Toggling the switch should cause the appropriate message to be written to your application's console.

The availability of the light is maintained automatically while your client is running but if you want to manually indicate that the light is unavailable, you can run:

```golang
l.SetAvailability(false)
```
