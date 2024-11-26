package quickzip

import (
	"bytes"
	"os"
	"runtime"
	"testing"
)

const (
	filesToZip = "./data/files"
	testFile   = "./data/test-file.tar.gz"
)

// hardcode pre-zipped file to check that our program zipped the test files correctly. This was zipped on my system using tar for thoroughness
var fileBytesWindows = []byte{80, 75, 3, 4, 20, 0, 8, 8, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 19, 0, 0, 0, 100, 97, 116, 97, 92, 102, 105, 108, 101, 115, 92, 116, 101, 115, 116, 46, 116, 120,
	116, 42, 73, 45, 46, 201, 204, 75, 87, 40, 46, 41, 77, 75, 43, 6, 4, 0, 0, 255, 255, 80, 75, 7, 8, 232,
	137, 191, 235, 20, 0, 0, 0, 14, 0, 0, 0, 80, 75, 1, 2, 20, 0, 20, 0, 8, 8, 8, 0, 0, 0, 0, 0, 232, 137,
	191, 235, 20, 0, 0, 0, 14, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 97, 116,
	97, 92, 102, 105, 108, 101, 115, 92, 116, 101, 115, 116, 46, 116, 120, 116, 80, 75, 5, 6, 0, 0, 0, 0, 1,
	0, 1, 0, 65, 0, 0, 0, 85, 0, 0, 0, 0, 0}

var fileBytesLinux = []byte{80, 75, 3, 4, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 100, 97, 116, 97, 47, 102, 105, 108, 101, 115, 47, 116, 101, 115, 116, 46, 116, 120, 116, 42,
	73, 45, 46, 201, 204, 75, 87, 40, 46, 41, 77, 75, 43, 6, 4, 0, 0, 255, 255, 80, 75, 7, 8, 232, 137, 191, 235,
	20, 0, 0, 0, 14, 0, 0, 0, 80, 75, 1, 2, 20, 0, 20, 0, 8, 0, 8, 0, 0, 0, 0, 0, 232, 137, 191, 235, 20, 0, 0,
	0, 14, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 97, 116, 97, 47, 102, 105, 108,
	101, 115, 47, 116, 101, 115, 116, 46, 116, 120, 116, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 65, 0, 0, 0, 85,
	0, 0, 0, 0, 0}

// the main testing function
func TestMain(t *testing.T) {
	/* setup */
	//create data dir
	err := os.Mkdir("data", 0777)
	if err != nil {
		t.Errorf("could not create data/ directory: %v", err)
	}
	//create nested dir for testing files
	err = os.Mkdir("data/files", 0777)
	if err != nil {
		t.Errorf("could not create data/files/ directory: %v", err)
	}
	//populate file to zip
	err = os.WriteFile("data/files/test.txt", []byte("testing stuffs"), 0777)
	if err != nil {
		t.Errorf("could not populate data/files/test.txt: %v", err)
	}

	//cleanup all new files and directories we created after test run
	t.Cleanup(func() {
		os.RemoveAll("data/")
	})

	//run actual zip test
	t.Run("Test zip", func(t *testing.T) {
		z := NewZipper(filesToZip)
		err := z.Zip(testFile)
		if err != nil {
			t.Errorf("could not zip file %s: %v", testFile, err)
		}
		defer os.Remove(testFile)

		ff, err := os.ReadFile(testFile)
		if err != nil {
			t.Errorf("could not read file %s: %v", testFile, err)
		}

		//os check
		if runtime.GOOS == "windows" {
			if !bytes.Equal(fileBytesWindows, ff) {
				t.Error("zipped files do not match")
			}
		} else {
			if !bytes.Equal(fileBytesLinux, ff) {
				t.Error("zipped files do not match")
			}
		}
	})
}
