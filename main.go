/*
 * @Author: Vincent Yang
 * @Date: 2025-08-19 18:15:50
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2025-08-19 18:23:37
 * @FilePath: /FixTwitter/main.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2025 by Vincent, All Rights Reserved.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Command line flags
	var service = flag.String("service", "fxtwitter.com", "Custom service domain to replace x.com with")
	flag.Parse()

	fmt.Println("Starting FixTwitter clipboard service...")
	fmt.Printf("Service will monitor clipboard and automatically replace x.com links with %s\n", *service)
	fmt.Println("Press Ctrl+C to stop the service")

	monitor := NewClipboardMonitor(*service)

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nStopping service...")
		os.Exit(0)
	}()

	// Start monitoring
	log.Printf("FixTwitter clipboard monitoring service started with replacement service: %s", *service)
	monitor.Start()
}
