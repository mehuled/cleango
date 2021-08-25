package main


type FileInformation struct {
	name string
	content []byte
	hash string
	size int64
	path string
	extension string
}

func getName(f FileInformation) string {
	return f.name
}

func getContent(f FileInformation) []byte {
	return f.content
}

func getExtension(f FileInformation) string {
	return f.extension
}