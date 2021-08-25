package main

import "time"

type FileInfo struct {
	name string
	lastAccessTime time.Time
}
func Push(info FileInfo) error {
	return nil
}

func Pop() (FileInfo ,error) {
	return FileInfo{name : "hello"}, nil
}