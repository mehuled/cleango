package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)


func traverseDir(hashes, files map[string]string, duplicates map[string]string, entries []os.FileInfo, directory string, extensions map[string]int32) {

	for _, entry := range entries{
		files[entry.Name()] = toReadableSize(entry.Size())
		err := Push(FileInfo{
			name:           entry.Name(),
			lastAccessTime: entry.ModTime(),
		})
		if err != nil {
			return
		}
		fullPath := path.Join(directory, entry.Name())

		if val, ok := extensions[path.Ext(entry.Name())]; ok {
			extensions[path.Ext(entry.Name())] = val + 1
		} else {
			extensions[path.Ext(entry.Name())] = 1
		}

		if !entry.Mode().IsDir() && !entry.Mode().IsRegular() {
			continue
		}

		if entry.IsDir() {
			continue
		}
		file, err := ioutil.ReadFile(fullPath)
		//fmt.Println(string(file))
		if err != nil {
			panic(err)
		}
		hash := sha1.New()
		if _, err := hash.Write(file); err != nil {
			panic(err)
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


/*
Bytes conversion constants
 */
const (
	tbBytes int64 =1000*1000*1000*1000
	gbBytes int64 =1000*1000*1000
	mbBytes int64 =1000*1000
	kbBytes int64 =1000
)
func toReadableSize(nbytes int64) string {
	if nbytes > tbBytes {
		return strconv.FormatInt(nbytes/(tbBytes), 10) + " TB"
	}
	if nbytes > gbBytes {
		return strconv.FormatInt(nbytes/(gbBytes), 10) + " GB"
	}
	if nbytes > mbBytes {
		return strconv.FormatInt(nbytes/(mbBytes), 10) + " MB"
	}
	if nbytes > kbBytes {
		return strconv.FormatInt(nbytes/kbBytes, 10) + " KB"
	}
	return strconv.FormatInt(nbytes, 10) + " B"
}


/*
deleteFile will take Full path of file as input and will delete that file.
 */
func deleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
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
	for key, val := range files {
		fmt.Printf("|name : %s |\t size : %s|\n", key, val)
	}

	fmt.Println("#Total duplicate files")
	fmt.Println(len(duplicates))

	fmt.Println("#Duplicate files")
	for _, val := range duplicates {
		fmt.Printf("%s\n", val)
	}

	fmt.Println("#Group by extensions")
	for key, val := range extensions {
		fmt.Printf("%s : %d\n", key, val)
	}

}
