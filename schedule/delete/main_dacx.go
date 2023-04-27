package main

import (
	"log"
)

/*
 */

func main() {
	log.Println("Deleting schedule", "ScheduleID", scheduleHandle.GetID())
	err = scheduleHandle.Delete(ctx)
	if err != nil {
		log.Fatalln("Unable to delete schedule", err)
	}
}