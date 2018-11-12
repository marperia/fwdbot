package main

import (
	"fmt"
	"github.com/Arman92/go-tdlib"
	"github.com/marperia/fwdbot/accounts"
	"github.com/marperia/fwdbot/menu"
	"os"
	"os/signal"
)

func main() {
	var err error
	tdlib.SetLogVerbosityLevel(1)
	tdlib.SetFilePath("./errors.txt")

	err = accounts.InitConfig()
	accounts.ReadConfigFile()
	if err != nil {
		fmt.Println("Can't initialise config:", err)
	}

	err = accounts.InitAccounts()
	accounts.ReadAccountsFile()
	if err != nil {
		fmt.Println("Can't initialise accounts:", err)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)

	for _, acc := range accounts.TdInstances {

		acc.LoginToTdlib()

		// Handle Ctrl+C
		go func(ac accounts.TdInstance) {
			<-c
			ac.TdlibClient.DestroyInstance()
			os.Exit(42)
		}(acc)
	}

	for {
		menu.CallMenu()
	}
}
