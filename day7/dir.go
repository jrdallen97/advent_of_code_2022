package main

import (
	"fmt"
	"sort"
	"strconv"
)

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:     name,
		parent:   parent,
		children: []*Dir{},
		files:    map[string]int{},
	}
}

type Dir struct {
	name     string
	children []*Dir
	parent   *Dir
	files    map[string]int
}

func (d *Dir) AddChild(name string) *Dir {
	// If the child already exists, don't add them
	child := d.GetChild(name)
	if child != nil {
		return child
	}
	// Otherwise, create, add and return them
	child = NewDir(name, d)
	d.children = append(d.children, child)
	return child
}

func (d *Dir) GetChild(name string) *Dir {
	for _, c := range d.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func (d *Dir) AddFile(sizeOrDir, file string) {
	// If it's a dir that's listed, add it
	if sizeOrDir == "dir" {
		d.AddChild(file)
		return
	}
	fmt.Println("adding file", file, "to", d.name)
	// Otherwise, work out its size and add it
	intSize, err := strconv.Atoi(sizeOrDir)
	if err != nil {
		panic(err)
	}
	d.files[file] = intSize
}

// List files in dir (in alphabetical order)
func (d *Dir) getFileNames() []string {
	names := make([]string, len(d.files))
	i := 0
	for name := range d.files {
		names[i] = name
		i++
	}
	sort.Strings(names)
	return names
}

// Tree pretty-prints the contents of the directory (recursively)
func (d *Dir) Tree() {
	d.tree(0)
}

func (d *Dir) tree(indent int) {
	fmt.Printf("%s- %s (Dir)\n", makeIndent(indent), d.name)
	for _, c := range d.children {
		c.tree(indent + 1)
	}
	for _, name := range d.getFileNames() {
		fmt.Printf("%s- %s (file, size=%d)\n", makeIndent(indent+1), name, d.files[name])
	}
}

// GetDirSizes gets flattened dir sizes (note: must have run CalcSize first).
// Also returns the total size of the current directory.
// This needs to be separate as the slice will contain double-counted files.
func (d *Dir) GetDirSizes() ([]*DirSize, int) {
	size := 0
	// Work out the size of this directory's files
	for _, s := range d.files {
		size += s
	}

	// Get the sizes of all this directory's children
	var sizes []*DirSize
	for _, c := range d.children {
		childSizes, childSize := c.GetDirSizes()
		size += childSize
		sizes = append(sizes, childSizes...)
	}

	return append(sizes, &DirSize{d.name, size}), size
}

type DirSize struct {
	name string
	size int
}

func makeIndent(levels int) string {
	indent := ""
	for i := 0; i < levels*2; i++ {
		indent += " "
	}
	return indent
}
