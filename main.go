package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strconv"
)

const Size = 1000
const Base = 10

func traverseDir(hashes, files map[string]string, duplicates map[string]string,  entries []os.FileInfo, directory string, extensions map[string]int32) {
	for _, entry := range entries {

		files[entry.Name()] = toReadableSize(entry.Size())
		fileInfo :=  FileInfo{}

		fileInfo.name = entry.Name()
		fileInfo.lastAccessTime = entry.ModTime()

		fullPath := path.Join(directory, entry.Name())

		if val, ok := extensions[path.Ext(entry.Name())]; ok {
			extensions[path.Ext(entry.Name())] = val+1
		} else {
			extensions[path.Ext(entry.Name())] = 1
		}

		if !entry.Mode().IsDir() && !entry.Mode().IsRegular() {
			continue
		}

		if entry.IsDir() {
			continue
		}
		file, err2 := ioutil.ReadFile(fullPath)

		if err2 != nil {
			panic(err2)
		}
		hash := sha1.New()
		if _, err3 := hash.Write(file); err3 != nil {
			panic(err3)
		}
		hashSum := hash.Sum(nil)
		hashString := fmt.Sprintf("%x", hashSum)
		if hashEntry, ok := hashes[hashString]; ok {
			duplicates[hashEntry] = fullPath
		} else {
			hashes[hashString] = fullPath
		}
	}
}

func toReadableSize(numberOfBytes int64) string {
	if numberOfBytes > int64(math.Pow(Size,4)) {
		return strconv.FormatInt(numberOfBytes/int64(math.Pow(Size,4)), Base) + " TB"
	}
	if numberOfBytes > int64(math.Pow(Size,3)) {
		return strconv.FormatInt(numberOfBytes/int64(math.Pow(Size,3)), Base) + " GB"
	}
	if numberOfBytes > int64(math.Pow(Size,2)) {
		return strconv.FormatInt(numberOfBytes/int64(math.Pow(Size,2)), Base) + " MB"
	}
	if numberOfBytes > Size {
		return strconv.FormatInt(numberOfBytes/Size, Base) + " KB"
	}
	return strconv.FormatInt(numberOfBytes, Base) + " B"
}

/**
1. Takes as input a command-line argument --dir which is an absolute path to a directory in the host filesystem.
2. Traverses over all the files in the given dir (excluding the hidden files and subdirectories).
3. Generates a summary of the directory, the summary contains the following:
    1. Name and size of all files in the directory.
    2. Name and count of duplicate files (if any).
    3. Count of files grouped by extension.
    4. (Bonus) List of 10 least recently opened files.
*/
func main() {
	var err error
	dir := flag.String("dir", "", "the dir to summarize")
	flag.Parse()

	if *dir == "" {
		panic("please provide a dir to summarize")
	}

	files := map[string]string{}
	hashes := map[string]string{}
	duplicates := map[string]string{}
	extensions := map[string]int32{}

	entries, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	traverseDir(hashes, files, duplicates, entries, *dir, extensions)

	fmt.Println("#File info")
	for fileName, fileSize := range files {
		fmt.Printf("|name : %s |\t size : %s|\n", fileName, fileSize)
	}

	fmt.Println("#Total duplicate files")
	fmt.Println(len(duplicates))

	fmt.Println("#Duplicate files")
	for _, val := range duplicates {
		fmt.Printf("%s\n",val)
	}

	fmt.Println("#Group by extensions")
	for key, val := range extensions {
		fmt.Printf("%s : %d\n",key, val)
	}

}
