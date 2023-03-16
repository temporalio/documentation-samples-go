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
	}
	w := worker.New(temporalClient, "fileprocessing", workerOptions)

	w.RegisterWorkflow(sessions.SomeFileProcessingWorkflow)
	w.RegisterActivity(&sessions.FileActivities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

/* @dacx
id: how-to-enable-sessions-on-a-worker
title: How to enable Sessions on a Worker
label: Sessions
description: Set EnableSessionWorker to true in the Worker options.
lines: 1-3, 7, 10-12, 22-34
@dacx */
