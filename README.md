# quick-zip

[![Pipeline](https://github.com/jake-young-dev/quick-zip/actions/workflows/pipeline.yaml/badge.svg?branch=master)](https://github.com/jake-young-dev/quick-zip/actions/workflows/pipeline.yaml)

A simple dependency-free Golang package for zipping system files

# usage

```
package main

import (
	"log"

	qz "github.com/jake-young-dev/quick-zip"
)

func main() {
	zppr, err := qz.NewZipper("dir/to/zip")
	if err != nil {
		panic(err)
	}
	ttz, err := zppr.Zip("newfile.tar.gz")
	if err != nil {
		panic(err)
	}

	//log time taken to zip
	log.Println(ttz)

	//newfile.tar.gz will be zipped with the contents from dir/to/zip
}
```

# changes
Breaking changes will not be introduced before v1.0.0 I am still working through the ideal usage
of this repo and things like the time benchmarks may be removed in the future.