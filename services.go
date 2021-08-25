package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path"
	"strconv"
)

func getExtensionsforfile(extensions map[string]int32,entry os.FileInfo){
	if val, ok := extensions[path.Ext(entry.Name())]; ok {
		extensions[path.Ext(entry.Name())] = val+1
	} else {
		extensions[path.Ext(entry.Name())] = 1
	}

}
func getDuplicatesforfile(file []byte, fullpath string, hashes, duplicates map[string]string){
	hash := sha1.New()
	if _, err := hash.Write(file); err != nil {
		panic(err)
	}
	hashSum := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashSum)
	if hashEntry, ok := hashes[hashString]; ok {
		duplicates[hashEntry] = fullpath
	} else {
		hashes[hashString] = fullpath
	}
}
func toReadableSize(nbytes int64) string {
	if nbytes > TB {
		return strconv.FormatInt(nbytes/TB, 10) + "TB"
	}
	if nbytes > GB {
		return strconv.FormatInt(nbytes/GB, 10) + " GB"
	}
	if nbytes > MB {
		return strconv.FormatInt(nbytes/MB, 10) + " MB"
	}
	if nbytes > KB {
		return strconv.FormatInt(nbytes/KB, 10) + " KB"
	}
	return strconv.FormatInt(nbytes, 10) + " B"
}