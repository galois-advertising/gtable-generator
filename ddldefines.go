package generator

import "log"

type DdlDefines struct {
	name, namespace string
	dataviews       map[string]string
}

func (gd *DdlDefines) load(file_name string) (ok bool) {
	if !FileExists(file_name) {
		log.Printf("File %s not exists.\n", file_name)
		return false
	}
	return true

}
