package quickzip

import (
	"bytes"
	"log"
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

var fileBytesLinux = []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 3, 237, 209, 75, 10, 194, 48, 20, 133, 225, 140, 93,
	69, 86, 208, 230, 182, 105, 178, 158, 128, 173, 84, 197, 129, 185, 5, 151, 175, 29, 8, 85, 124, 32, 180,
	136, 248, 127, 147, 51, 200, 133, 28, 56, 69, 185, 78, 154, 202, 174, 223, 183, 185, 52, 203, 112, 23, 49,
	54, 99, 74, 108, 220, 52, 175, 140, 248, 88, 137, 84, 151, 136, 198, 137, 4, 239, 141, 109, 22, 234, 115,
	99, 200, 154, 142, 214, 154, 109, 218, 181, 175, 238, 222, 189, 255, 168, 98, 186, 191, 182, 89, 11, 61,
	233, 204, 127, 140, 3, 135, 224, 159, 239, 47, 241, 110, 255, 186, 14, 149, 177, 110, 230, 30, 15, 253, 249,
	254, 227, 228, 253, 97, 99, 179, 14, 93, 151, 87, 223, 174, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 224, 3,
	103, 19, 24, 237, 5, 0, 40, 0, 0}

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

		log.Println(runtime.GOOS)
		log.Println(ff)
		log.Println(fileBytesLinux)
		//is this local or pipeline?
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
