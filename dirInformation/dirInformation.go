package dirInformation

type DirInformation struct {
	Files map[string]string
	Hashes map[string]string
	Duplicates map[string]string
	Extensions map[string]int32
}

func (dir *DirInformation) InitializeDirInformation(){
	dir.Files=make(map[string]string)
	dir.Hashes=make(map[string]string)
	dir.Duplicates=make(map[string]string)
	dir.Extensions=make(map[string]int32)
}