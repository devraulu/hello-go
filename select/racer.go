package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} { // why chan struct{} and not bool? struct{} is the smallest data type available from memory perspective, no alloc required
	// always use make to create a channel
	// if you use var ch chan struct{} it will be init to the type's default value which is nil in the case of chan
	// trying to send a nil channel will block forever
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
