# quick-zip

[![Pipeline](https://github.com/jake-young-dev/quick-zip/actions/workflows/pipeline.yaml/badge.svg?branch=master)](https://github.com/jake-young-dev/quick-zip/actions/workflows/pipeline.yaml)

A simple Golang package for compressing directories into .tar.gz files

# usage

```
package main

import (
	"log"

	qz "github.com/jake-young-dev/quick-zip"
)

func main() {
	z, err := qz.NewZipper("dir/to/zip")
	if err != nil {
		panic(err)
	}
	err = z.Zip("newfile.tar.gz")
	if err != nil {
		panic(err)
	}

	//newfile.tar.gz will contain the contents from dir/to/zip
}
```
