package cron

import (
	"time"

	"github.com/robfig/cron"
)

func FirstCron() {
	c := cron.New()

	// Define the Cron job schedule
	c.AddFunc("* * * * *", func() {
		SendMail() //call from emailer.go
	})

	// Start the Cron job scheduler
	c.Start()

	// Wait for the Cron job to run
	time.Sleep(5 * time.Minute)

	// Stop the Cron job scheduler
	c.Stop()
}
