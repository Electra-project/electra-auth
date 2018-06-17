// Adapted from https://gist.github.com/albulescu/e61979cc852e4ee8f49c.

package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

func download(url string, dest string, size int64) {
	file := path.Base(url)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(file)

	out, err := os.Create(path.String())

	if err != nil {
		fmt.Println(path.String())
		panic(err)
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	done := make(chan int64)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	done <- n
}

// DownloadBinary downloads the Electra daemon binary.
func DownloadBinary() {
	var fileName string
	var size int64
	if runtime.GOOS == "windows" {
		fileName = "electrad-win32-x64.exe"
		size = 7167488
	} else if runtime.GOOS == "darwin" {
		fileName = "electrad-darwin-x64"
		size = 18828580
	} else {
		fileName = "electrad-linux-x64"
		size = 73570640
	}

	uriBase := "https://github.com"
	uriPath := "/Electra-project/electra/releases/download/v1.2.0-beta.2/"

	download(uriBase+uriPath+fileName, "./bin", size)
}
