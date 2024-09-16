package reload

import (
	"flag"
	"fmt"
	"os"
)

var reloadArg string

func init() {
	flag.StringVar(&reloadArg, "s", "reload", "Send signal to process. Supported: reload.")
}

func CheckArg(sig os.Signal, pidfile string) bool {
	if len(reloadArg) <= 0 {
		return false
	}

	switch reloadArg {
	case "reload":
		if err := SendSignal(pidfile, sig); err != nil {
			fmt.Println("Send signal error:", err)
		}
	default:
		fmt.Println("Unsupported signal:", reloadArg)
	}
	return true
}

// using a specific SIG to reload something
type Reloader func() error

// Setup reload signal monitor
// Calling `reloader` when `sig` triggerred
func Setup(sig os.Signal, pid string, reloader Reloader) {
	chSig := InstallSignal()
	go func() {
		for receivedSig := range chSig {
			if receivedSig == sig {
				if err := reloader(); err != nil {
					fmt.Println("reload error:", err)
				}
			}
		}
	}()
}
