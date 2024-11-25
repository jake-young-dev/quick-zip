package quickzip

import (
	"bytes"
	"os"
	"testing"
)

const (
	filesToZip  = "./data/files"
	testFile    = "./data/test-file.tar.gz"
	compareFile = "./data/compare-file.tar.gz"
)

// tests the main zip functionality, requires a /data/ directory in the project root containing a files directory and
// a already zipped copy of the files directory for comparison.
func TestZip(t *testing.T) {
	z := NewZipper(filesToZip)
	err := z.Zip(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)

	f, err := os.ReadFile(compareFile)
	if err != nil {
		t.Fatal(err)
	}

	ff, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(f, ff) {
		t.Fatal("files do not match")
	}
}
