package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "check_vitelity_balance 0.1"
const api_url = "https://api.vitelity.net/api.php"

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [args]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

var verbose int

func main() {

	verbose = 0
	warn_level := 10.00
	crit_level := 5.00
	username := ""
	password := ""

	flag.Usage = usage
	flag.IntVar(&verbose, "verbose", 0, "verbosity level (0-3)")
	flag.Float64Var(&warn_level, "warning", 100.00, "warn when balance is below this value")
	flag.Float64Var(&crit_level, "critical", 25.00, "critical when balance is below this value")
	flag.StringVar(&username, "username", "", "API Username")
	flag.StringVar(&password, "password", "", "API Password")
	flag.Parse()

	if verbose >= 3 {
		fmt.Println("Verbosity level:", verbose)
		fmt.Println("Warning level:", warn_level)
		fmt.Println("Critical level:", crit_level)
	}

	balance, err := get_balance(username, password)
	if err != nil {
		fmt.Println("Balance Unknown:", err)
		os.Exit(3)
	}

	if verbose >= 2 {
		fmt.Println("Balance:", balance)
	}

	// check if balance is below critical threshold
	if balance <= crit_level {
		fmt.Printf("Balance Critical - $%.2f\n", balance)
		os.Exit(2)
	}

	// check if balance is below warning threshold
	if balance <= warn_level {
		fmt.Printf("Balance Warning - $%.2f\n", balance)
		os.Exit(1)
	}

	// balance must be ok
	fmt.Printf("Balance OK - $%.2f\n", balance)
	os.Exit(0)
}
