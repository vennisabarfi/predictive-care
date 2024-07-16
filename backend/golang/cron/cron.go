package cron

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
)

// adding Mutex to sync and avoid race conditions

func NewsletterCron() {

	c := cron.New()

	/* Job schedule: Send proverbs every day to user*/
	err := c.AddFunc("@daily", func() {
		SendMail() //call from emailer.go
	})
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	// Start the Cron job scheduler
	c.Start()
	fmt.Println("Cron job has started. Sending proverb to user email!")

	// Shutdown program gracefully
	sigChan := make(chan os.Signal, 1)                    //if shutdown signal received
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM) //gracefully stop job scheduler, db process and log event
	<-sigChan

	fmt.Println("Gracefully shutting down cron job scheduler...")
	c.Stop()
	fmt.Println("Cron job stopped!")

}
