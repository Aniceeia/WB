package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var (
	ErrServerFailed   = "error from %s: %v\n"
	ErrAllServersDown = "all servers unavailable: %v\n"
)

const TimeFormat = "2006-01-02 15:04:05 MST"

func main() {
	servers := []string{
		"time.google.com",
		"pool.ntp.org",
		"time.windows.com",
		"ntp1.stratum1.ru",
	}

	var ntpTime time.Time
	var err error

	for _, server := range servers {
		ntpTime, err = ntp.Time(server)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, ErrServerFailed, server, err)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, ErrAllServersDown, err)
		os.Exit(1)
	}

	fmt.Println(ntpTime.Format(TimeFormat))
}
