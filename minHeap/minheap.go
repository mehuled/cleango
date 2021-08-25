package main

import "time"

type FileInfoaccess struct {
	Name string
	Last_access_time time.Time
}
func Push(info FileInfoaccess) error {
	return nil
}

func Pop() (FileInfoaccess ,error) {
	return FileInfoaccess{Name : "hello"}, nil

}