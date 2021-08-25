package main

import "time"

type FileInfoaccess struct {
	name string
	last_access_time time.Time
}
func Push(info FileInfoaccess) error {
	return nil
}

func Pop() (FileInfoaccess ,error) {
	return FileInfoaccess{name : "hello"}, nil

}