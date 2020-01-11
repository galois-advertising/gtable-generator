package generator

import (
	"testing"
)

func TestArchiveFile(t *testing.T) {
	ArchiveFile("https://github.com/galois-advertising/ghead/blob/master/README.md", "./")
}
