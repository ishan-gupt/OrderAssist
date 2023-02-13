package utils

import "go.mau.fi/whatsmeow/store"

func NameChange(names string) {
	var v = [3]uint32{0, 1, 2}
	store.SetOSInfo(names, v)
}
