# `tantalic.com/anybar`

A [go][golang] package for interacting with [AnyBar][anybar], a macOS menu bar status indicator application.

The package makes it simple to set the AnyBar icon style:

```go
package main

import (
    "tantalic.com/anybar"
)

func main() {
    client := anybar.Client{}
    client.Set(anybar.Green)
}
```


## Installation
```shell
go get -u tantalic.com/anybar
```


## API
- [type anybar.Client](#anybarclient)
    + [`Set(s Style)`](#sets-anybarstyle)
    + [`Start()`](#start)
    + [`Quit()`](#quit)
- [type anybar.Style](#anybarstyle)


### `anybar.Client`

A client is used to send commands to an instance of the AnyBar app.

``` go
type Client struct {
    Port     int
    Hostname string
}
```

`Client.Port` is the UDP port used to control the AnyBar application. Default value: `1738`.
`Client.Hostname` is the host where the AnyBar app is running. Default value `localhost`.


#### `Set(s anybar.Style)`

``` go
func (c *Client) Set(s anybar.Style) error
```

Sets the style of the AnyBar menu bar icon.


#### `Start()`

``` go
func (c *Client) Start() error
```

Opens the AnyBar app. If an instance of AnyBar is already running on `Client.Port` that instance will be quit first. This function only works on macOS and requires AnyBar.app to be installed.


#### `Quit()`
``` go
func (c *Client) Quit() error
```

Send the command to exit the AnyBar app.


### `anybar.Style`

A style represents one of the available appearances that can be set to the AnyBar icon in the menu bar. You should not create your own instances but instead use on of following constants:

- `anybar.White`
- `anybar.Red`
- `anybar.Orange`
- `anybar.Yellow`
- `anybar.Green`
- `anybar.Cyan`
- `anybar.Blue`
- `anybar.Purple`
- `anybar.Black`
- `anybar.Question`
- `anybar.Exclamation`



[golang]: http://golang.org
[anybar]: https://github.com/tonsky/AnyBar
