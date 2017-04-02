package anybar_test

import "tantalic.com/anybar"
import "time"

// A client is used to set the style of the AnyBar icon in the menu bar:
func ExampleClient() {
	var client anybar.Client
	client.Set(anybar.Green)
}

// The Port and Hostname fields control what AnyBar server the client
// communicates with:
func ExampleClient_options() {
	client := anybar.Client{
		Port:     1738,
		Hostname: "Someones-Mac.local",
	}
	client.Set(anybar.Blue)
}

// The following example uses time.Tick to blink the AnyBar menu bar icon
func ExampleClient_ticker() {
	ticker := time.Tick(700 * time.Millisecond)

	var client anybar.Client
	for {
		client.Set(anybar.Red)
		<-ticker
		client.Set(anybar.Green)
		<-ticker
		client.Set(anybar.White)
		<-ticker
	}
}
