package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create Schedule

}

/* @dacx
id: how-to-describe-a-schedule-in-go
title: How to describe a Schedule in Go
label: Describe Schedule
description: Describe a Schedule in Go.
lines:
@dacx */