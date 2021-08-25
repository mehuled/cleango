package Services

import "time"

type FileInfo struct {
	name string
	last_access_time time.Time
}
func Push(info FileInfo) error {

	return nil
}

func Pop() (FileInfo ,error) {
	return FileInfo{name : "hello"}, nil

}