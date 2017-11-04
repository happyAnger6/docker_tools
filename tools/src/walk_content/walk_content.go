package main

import (
	"io/ioutil"
	"path/filepath"
	"github.com/go-digest"
	"fmt"
	"os"
)

const (
	contentRootDir="/var/lib/docker/image/overlay2/imagedb/"
	contentDir="content"
	metadataDir="metadata"
)
func main(){
	dir, err := ioutil.ReadDir(filepath.Join(contentRootDir, contentDir, string(digest.Canonical)))
	if err != nil {
		fmt.Printf("open %s err: %v\r\n", contentDir, err)
		os.Exit(-1)
	}



}
