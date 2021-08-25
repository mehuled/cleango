package main

import (
	"flag"
	"io/ioutil"
)


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

	traverseDirectory(hashes, files, duplicates, entries, *dir, extensions)

	displayFileinformation(extensions,files,duplicates)


}

