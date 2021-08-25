package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func traverseDirectory(hashes, files map[string]string, duplicates map[string]string,  entries []os.FileInfo, directory string, extensions map[string]int32) {
	for _, entry := range entries {

		files[entry.Name()] = toReadableSize(entry.Size())
		err := Push(FileInfoaccess{
			name: entry.Name(),
			last_access_time: entry.ModTime(),
		})
		if err != nil {
			return
		}
		fullpath := path.Join(directory, entry.Name())

		//get the extensions for a particular entry
		getExtensionsforfile(extensions,entry)

		if !entry.Mode().IsDir() && !entry.Mode().IsRegular() {
			continue
		}

		if entry.IsDir() {
			continue
		}
		file, err := ioutil.ReadFile(fullpath)

		if err != nil {
			panic(err)
		}
		//get the duplicates for a particular file
		getDuplicatesforfile(file,fullpath,hashes,duplicates)

	}
}


func deleteFile(file_path string) error {
	err := os.Remove(file_path)
	if err != nil {
		return err
	}
	return nil
}

func displayFileinformation(extensions map[string]int32,files,duplicates map[string]string){
	fmt.Println("#File info")
	for key, val := range files {
		fmt.Printf("|name : %s |\t size : %s|\n",key,val)
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
