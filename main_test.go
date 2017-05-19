package main

import (
	"testing"
)

func TestNewCredentials(t *testing.T) {
	ips, channels, chains := setUp("db.txt","ip_test.txt", "html")

	launchChains(ips, channels, chains)

	waitForResponse(channels)
}