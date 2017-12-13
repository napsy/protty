package main

import (
	gkeyring "github.com/rakyll/keys"
)

type keyring interface {
	Get(service, username string) (string, error)
	Set(service, username, password string) error
	Delete(service, username string) error
}

var keys keyring

func init() {
	// Use gnome-keyring by default
	k, err := gkeyring.New()
	if err != nil {
		panic(err)
	}
	keys = k

}
