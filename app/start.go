package app

import (
	"fwdbot/accounts"
	"github.com/Arman92/go-tdlib"
	"time"
)

func Start() {
	// initialise TdInstances var of accounts package
	accounts.ReadAccountsFile()
	// initialise Configs var of config package
	accounts.ReadConfigFile()
	for _, acc := range accounts.TdInstances {
		go func(funcacc accounts.TdInstance) {
			funcacc.LoginToTdlib()

			receiver := funcacc.TdlibClient.AddEventReceiver(&tdlib.UpdateNewMessage{}, accounts.MessageFilter, 10)
			for newMsg := range receiver.Chan {
				accounts.NewMessageHandle(newMsg, funcacc)
			}

			accounts.CreateUpdateChannel(funcacc.TdlibClient)
		}(acc)
		time.Sleep(300 * time.Millisecond)
	}
}
