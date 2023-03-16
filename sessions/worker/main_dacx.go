package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"documentation-samples-go/sessions"
)

/*
Set `EnableSessionWorker` to `true` in the Worker options.
*/

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()
	// Enable Sessions for this Worker.
	workerOptions := worker.Options{
		EnableSessionWorker: true,
		// This configures the maximum allowed concurrent sessions
		// Only customize this value if you need to.
		MaxConcurrentSessionExecutionSize: 1000,
	}
	w := worker.New(temporalClient, "fileprocessing", workerOptions)
	w.RegisterWorkflow(sessions.SomeFileProcessingWorkflow)
	w.RegisterActivity(&sessions.FileActivities{})
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

/*
To limit the number of concurrent Sessions running on a Worker, set the `MaxConcurrentSessionExecutionSize` field of `worker.Options` to the desired value.
By default, this field is set to a very large value, so there's no need to manually set it if no limitation is needed.

If a Worker hits this limitation, it won't accept any new `CreateSession()` requests until one of the existing sessions is completed. `CreateSession()` will return an error if the session can't be created within `CreationTimeout`.
*/

/* @dacx
id: how-to-enable-sessions-on-a-worker
title: How to enable Sessions on a Worker
label: Enable Sessions
description: Set EnableSessionWorker to true in the Worker options.
lines: 12-14, 16, 25-27, 31-35, 39
@dacx */

/* @dacx
id: how-to-configure-max-concurrent-sessions
title: How to configure the maximum concurrent Sessions on the Worker
label: Max concurrent Sessions
description: Set EnableSessionWorker to true in the Worker options.
lines: 16, 26, 28-30, 39, 41-46
@dacx */