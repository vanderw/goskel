package reload

import (
	"os"
	"os/signal"
)

func InstallSignal(sigs ...os.Signal) chan os.Signal {
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, sigs...)
	return chSig
}
