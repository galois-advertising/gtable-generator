package main

import "log"

type GqlDefines struct {
	name string
}

func (gd *GqlDefines) load(file_name string) (ok bool) {
	if !FileExists(file_name) {
		log.Printf("File %s not exists.\n", file_name)
		return false
	}
	return true
}
