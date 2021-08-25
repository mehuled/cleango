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

func traverseDir(hashes, files map[string]string, duplicates map[string]string,  entries []os.FileInfo, directory string, extensions map[string]int32) {
	for _, entry := range entries {

		files[entry.Name()] = toReadableSize(entry.Size())
		err := Push(FileInfo{
			name: entry.Name(),
			last_access_time: entry.ModTime(),
		})
		if err != nil {
			return
		}
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

func toReadableSize(numOfBytes int64) string {


	if numOfBytes > 1000*1000*1000*1000 {
		return strconv.FormatInt(numOfBytes/(1000*1000*1000*1000), 10) + " TB"
	}
	if numOfBytes > 1000*1000*1000 {
		return strconv.FormatInt(numOfBytes/(1000*1000*1000), 10) + " GB"
	}
	if numOfBytes > 1000*1000 {
		return strconv.FormatInt(numOfBytes/(1000*1000), 10) + " MB"
	}
	if numOfBytes > 1000 {
		return strconv.FormatInt(numOfBytes/1000, 10) + " KB"
	}
	return strconv.FormatInt(numOfBytes, 10) + " B"
}

func deleteFile(file_path string) error {
	err := os.Remove(file_path)
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
	dir := flag.String("dir", "", "the directory to summarize")
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

