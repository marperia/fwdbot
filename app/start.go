package app

import (
	"github.com/Arman92/go-tdlib"
	"github.com/marperia/fwdbot/accounts"
	"time"
)

func HandleMessages() {
	for i := range accounts.TdInstances {
		go func() {
			accounts.TdInstances[i].LoginToTdlib()

			receiver := accounts.TdInstances[i].TdlibClient.AddEventReceiver(&tdlib.UpdateNewMessage{}, accounts.MessageFilter, 10)
			for newMsg := range receiver.Chan {
				accounts.NewMessageHandle(newMsg, accounts.TdInstances[i])
			}

			accounts.CreateUpdateChannel(accounts.TdInstances[i].TdlibClient)
		}()
		time.Sleep(300 * time.Millisecond)
	}
}

func Start() {
	// initialise TdInstances var of accounts package
	accounts.ReadAccountsFile()
	// initialise Configs var of config package
	accounts.ReadConfigFile()

	HandleMessages()
}
