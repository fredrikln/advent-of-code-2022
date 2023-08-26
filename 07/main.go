package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	fs := NewFileSystem()

	// input := parsing.ReadRowsToStringSlice("input_example.txt")
	input := parsing.ReadRowsToStringSlice("input.txt")

	for _, line := range input {
		fs.ParseLine(line)
	}

	dirs := Find100kDirs(fs.Root)

	var totalSize int
	for _, dir := range dirs {
		totalSize += dir.GetSize()
	}

	fmt.Println("Part 1:", totalSize)

	diskSize := 70000000
	filesSize := fs.Root.GetSize()
	requiredSize := 30000000

	freeSpace := diskSize - filesSize
	sizeToRemove := requiredSize - freeSpace

	possibleTargets := []*Dir{fs.Root}
	possibleTargets = append(possibleTargets, FindDirsToDelete(fs.Root, sizeToRemove)...)

	smallestSize := diskSize
	var smallestDir *Dir

	for _, dir := range possibleTargets {
		if dir.GetSize() < smallestSize {
			smallestDir = dir
			smallestSize = dir.GetSize()
		}
	}

	fmt.Println("Part 2:", smallestDir.GetSize())
}

func Find100kDirs(root *Dir) []*Dir {
	var dirs []*Dir

	for _, dir := range root.Directories {
		if dir.GetSize() <= 100000 {
			dirs = append(dirs, dir)
		}

		dirs = append(dirs, Find100kDirs(dir)...)
	}

	return dirs
}

func FindDirsToDelete(root *Dir, sizeRequired int) []*Dir {
	var dirs []*Dir

	for _, dir := range root.Directories {
		if dir.GetSize() >= sizeRequired {
			dirs = append(dirs, dir)
		}

		dirs = append(dirs, FindDirsToDelete(dir, sizeRequired)...)
	}

	return dirs
}

type File struct {
	Name string
	Size int
}

func (f *File) GetSize() int {
	return f.Size
}

func (f *File) String() string {
	return fmt.Sprint("- ", f.Name, " (file, size=", f.Size, ")")
}

type Dir struct {
	Name        string
	ParentDir   *Dir
	Directories []*Dir
	Files       []*File
	Level       int
	Size        *int
}

func (d *Dir) AddFile(file *File) {
	d.Files = append(d.Files, file)
}

func (d *Dir) AddDir(dir *Dir) {
	d.Directories = append(d.Directories, dir)
}

func (d *Dir) GetSize() int {
	if d.Size != nil {
		return *d.Size
	}

	size := 0

	for _, dir := range d.Directories {
		size += dir.GetSize()
	}

	for _, file := range d.Files {
		size += file.GetSize()
	}

	d.Size = &size

	return size
}

func (d *Dir) String() string {
	out := fmt.Sprint("- ", d.Name, " (dir, size=", d.GetSize(), ")\n")

	for _, dir := range d.Directories {
		out = fmt.Sprint(out, strings.Repeat("  ", d.Level), dir, "\n")
	}

	for _, file := range d.Files {
		out = fmt.Sprint(out, strings.Repeat("  ", d.Level+1), file, "\n")
	}

	return out
}

type FileSystem struct {
	Root   *Dir
	CurDir *Dir
}

func NewFileSystem() *FileSystem {
	root := &Dir{
		ParentDir:   nil,
		Name:        "/",
		Directories: make([]*Dir, 0),
		Files:       make([]*File, 0),
		Level:       1,
	}

	return &FileSystem{
		Root:   root,
		CurDir: root,
	}
}

func (fs *FileSystem) ParseLine(line string) {
	parts := strings.Split(line, " ")

	if parts[0] == "$" {
		command := parts[1]

		if command == "ls" {
			return
		}

		if command == "cd" {
			dirname := parts[2]

			if dirname == ".." {
				fs.CurDir = fs.CurDir.ParentDir
				return
			}

			if dirname == "/" {
				fs.CurDir = fs.Root
				return
			}

			for _, dir := range fs.CurDir.Directories {
				if dir.Name == dirname {
					fs.CurDir = dir

					return
				}
			}

			panic("Could not find directory")
		}

		panic("Unknown command")
	}

	if parts[0] == "dir" {
		dirname := parts[1]

		dir := &Dir{
			Name:        dirname,
			ParentDir:   fs.CurDir,
			Directories: make([]*Dir, 0),
			Files:       make([]*File, 0),
			Level:       fs.CurDir.Level + 1,
		}

		fs.CurDir.AddDir(dir)

		return
	}

	filesize, err := strconv.Atoi(parts[0])

	if err != nil {
		panic("Invalid filesize")
	}

	filename := parts[1]

	file := &File{
		Name: filename,
		Size: filesize,
	}

	fs.CurDir.AddFile(file)
}
