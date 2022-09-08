package util

import (
	"io/fs"
	"os"
)

type Iterator struct {
	root  string
	index int
	files []fs.DirEntry
	subs  map[fs.DirEntry]Iterator
}

func GetIterator(root string) Iterator {
	return Iterator{root, 0, nil, make(map[fs.DirEntry]Iterator)}
}

func (i *Iterator) ReadDir() {
	if i.files != nil {
		return
	}
	files, err := os.ReadDir(i.root)
	if err != nil {
		panic(err)
	}
	i.files = files
}

func (i *Iterator) Next() string {
	i.ReadDir()
	if i.index >= len(i.files) {
		return ""
	}
	file := i.files[i.index]
	if !file.IsDir() {
		i.index += 1
		return i.root + "/" + file.Name()
	}
	if _, ok := i.subs[file]; !ok {
		i.subs[file] = GetIterator(i.root + "/" + file.Name())
	}
	if i.subs[file].HasNext() {
		sub := i.subs[file]
		next := sub.Next()
		i.subs[file] = sub
		return next
	}
	delete(i.subs, file)
	i.index++
	return i.Next()
}

func (i Iterator) HasNext() bool {
	i.ReadDir()
	return i.index < len(i.files)
}
