package main

import (
	"container/list"
	"fmt"
)

type session interface {
	typ() string
	name() string
	toMap() map[string]string
}

type sshConnection struct {
	sessionName string
	sshHost     string
	user        string
}

var sessions *list.List

func sshSessionFromMap(name string, session map[string]string) (*sshConnection, error) {
	return newSSHSession(name, session["host"], session["user"])
}

func (s *sshConnection) toMap() map[string]string {
	return map[string]string{
		"host": s.sshHost,
		"user": s.user,
	}
}

func (s *sshConnection) name() string {
	return s.sessionName
}

func newSSHSession(name, host, user string) (*sshConnection, error) {
	key, err := keys.Get("protty", "ssh-"+host+"user-"+user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KEY: %v\n", key)
	return &sshConnection{}, nil
}
