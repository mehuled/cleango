package main


type FileInfo2 struct {
	name string
	content []byte
	hash string
	size int64
	path string
	extension string
}

func getName(f FileInfo2) string {
	return f.name
}

func getContent(f FileInfo2) []byte {
	return f.content
}

func getExtension(f FileInfo2) string {
	return f.extension
}