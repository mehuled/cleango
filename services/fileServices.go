package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/razorpay/clean-go/constants"
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

func getDuplicatesforfile(file []byte, filepath string, hashes, duplicates map[string]string){
	hash := sha1.New()
	if _, err := hash.Write(file); err != nil {
		panic(err)
	}
	hashSum := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashSum)
	if hashEntry, ok := hashes[hashString]; ok {
		duplicates[hashEntry] = filepath
	} else {
		hashes[hashString] = filepath
	}
}
func toReadableSize(nbytes int64) string {
	if nbytes > constants.TB {
		return strconv.FormatInt(nbytes/constants.TB, 10) + "TB"
	}
	if nbytes > constants.GB {
		return strconv.FormatInt(nbytes/constants.GB, 10) + " GB"
	}
	if nbytes > constants.MB {
		return strconv.FormatInt(nbytes/constants.MB, 10) + " MB"
	}
	if nbytes > constants.KB {
		return strconv.FormatInt(nbytes/constants.KB, 10) + " KB"
	}
	return strconv.FormatInt(nbytes, 10) + " B"
}


