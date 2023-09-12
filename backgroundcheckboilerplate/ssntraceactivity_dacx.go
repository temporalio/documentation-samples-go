package backgroundcheckboilerplate

import (
	"context"
)

// SSNTraceActivity is your custom Activity Definition.
func SSNTraceActivity(ctx context.Context, param string) (*string, error) {
	// This is where a call to another service is made
	// Here we are pretending that the service that does SSNTrace returned "pass"
	result := "pass"
	return &result, nil
}

/* @dacx
id: backgroundcheck-boilerplate-ssntrace
title: How to develop an Activity Definition in Go
label: Boilerplate SSNTrace Activity Definition
description: In the Temporal Go SDK programming model, an Activity Definition is an exportable function or a `struct` method.
tags:
- go sdk
- code sample
- activity
lines: 1-7, 37-56, 70-82
@dacx */
