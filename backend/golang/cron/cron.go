package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func emailSend() {

}

func FirstCron() {
	c := cron.New()

	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		fmt.Println("Hello world!")
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}
