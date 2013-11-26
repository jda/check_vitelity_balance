package main

import (
	"flag"
	"fmt"
	//"github.com/jda/go-nagios/nagios"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [args]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func get_balance(user string, pass string, timeout int) (int64, error) {
	return 0, nil
}

func main() {
	verbose := 0
	warn_level := 10000
	crit_level := 500
	timeout := 10
	username := ""
	password := ""

	flag.Usage = usage
	flag.IntVar(&verbose, "verbose", 0, "verbosity level (0-3)")
	flag.IntVar(&warn_level, "warning", 1000, "warn when balance is below this value")
	flag.IntVar(&crit_level, "critical", 500, "critical when balance is below this value")
	flag.IntVar(&timeout, "timeout", 10, "API connection timeout in seconds")
	flag.StringVar(&username, "username", "", "API Username")
	flag.StringVar(&password, "password", "", "API Password")
	flag.Parse()

	if verbose >= 3 {
		fmt.Println("Verbosity level:", verbose)
		fmt.Println("Warning level:", warn_level)
		fmt.Println("Critical level:", crit_level)
		fmt.Println("Request timeout:", timeout)
	}

	balance, err := get_balance(username, password, timeout)
	if err != nil {
		fmt.Println("Balance Unknown:", err)
		os.Exit(3)
	}

	if verbose >= 2 {
		fmt.Println("Balance:", balance)
	}
}
