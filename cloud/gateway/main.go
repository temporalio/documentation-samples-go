package main

import (
	"go.temporal.io/sdk/client"
)

/*
Use the [`Dial()`](https://pkg.go.dev/go.temporal.io/sdk/client#Dial) API available in the [`go.temporal.io/sdk/client`](https://pkg.go.dev/go.temporal.io/sdk/client) package to create a new [`Client`](https://pkg.go.dev/go.temporal.io/sdk/client#Client).

Provide [`HostPort`](https://pkg.go.dev/go.temporal.io/sdk/internal#ClientOptions).

Set a custom Namespace name in the Namespace field on an instance of the Client Options.

Use the [`ConnectionOptions`](https://pkg.go.dev/go.temporal.io/sdk/client#ConnectionOptions) API to set mTLS.
*/

func main() {
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return err
	}
	client, err := client.Dial(client.Options{
		HostPort:  "your-custom-namespace.tmprl.cloud:7233",
		Namespace: "your-custom-namespace",
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{Certificates: []tls.Certificate{cert}},
		},
	}
	defer temporalClient.Close()
	// ...
}

