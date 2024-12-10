package day9

import (
	"container/list"
	"fmt"
	"slices"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetString(2024, 9)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

type File struct {
	Id     int
	Index  int
	Length int
	IsFile bool
}

func part1(input string) (result int) {
	isFile := true
	id, index := 0, 0
	disk := list.New()
	for _, d := range strings.Split(input, "") {
		digit := utils.MustAtoi(d)
		if isFile {
			disk.PushBack(File{Id: id, Index: index, Length: digit, IsFile: true})
			id++
		}
		index += digit
		isFile = !isFile
	}
	diskSlice := diskToArray(disk)
	for i := 0; i < len(diskSlice); i++ {
		if diskSlice[i] == -1 {
			for diskSlice[len(diskSlice)-1] == -1 {
				diskSlice = deleteLast(diskSlice)
			}
			if i >= len(diskSlice) {
				break
			}
			diskSlice[i] = diskSlice[len(diskSlice)-1]
			diskSlice = deleteLast(diskSlice)
		}
	}

	result = checksum(diskSlice)
	return
}

func deleteLast[S ~[]E, E any](slice S) S {
	return slices.Delete(slice, len(slice)-1, len(slice))
}

func checksum(disk []int) (sum int) {
	for i, v := range disk {
		if v == -1 {
			continue
		}
		sum += i * v
	}
	return
}

func diskToArray(disk *list.List) []int {
	cursor := disk.Front()
	cursorFile := cursor.Value.(File)
	diskArray := []int{}
	for i := 0; ; i++ {
		if i >= cursorFile.Index && i < (cursorFile.Index+cursorFile.Length) {
			if cursorFile.IsFile {
				diskArray = append(diskArray, cursorFile.Id)
			} else {
				diskArray = append(diskArray, -1)
			}
		} else {
			diskArray = append(diskArray, -1)
		}
		if (cursorFile.Index + cursorFile.Length - 1) == i {
			cursor = cursor.Next()
			if cursor == nil {
				return diskArray
			}
			cursorFile = cursor.Value.(File)
		}
	}
}

func part2(input string) (result int) {
	fileId, idx := 0, 0
	disk := list.New()
	for i, d := range strings.Split(input, "") {
		digit := utils.MustAtoi(d)
		disk.PushBack(File{
			Id:     fileId,
			Index:  idx,
			Length: digit,
			IsFile: i%2 == 0,
		})
		idx += digit
		if i%2 == 0 {
			fileId++
		}
	}
	fileId--
	file, fileData := elementAndFile(disk.Back())
	for fileId >= 0 {
		if fileData.Id != fileId || !fileData.IsFile {
			file, fileData = elementAndFile(file.Prev())
			continue
		}
		for space, spaceData := elementAndFile(disk.Front()); space != nil; space, spaceData = elementAndFile(space.Next()) {
			if spaceData.IsFile {
				continue
			}
			if spaceData.Length < fileData.Length {
				continue
			}
			if spaceData.Id >= (fileData.Id + 1) {
				continue
			}
			spaceData.Length -= fileData.Length
			space.Value = spaceData
			disk.InsertBefore(fileData, space)

			if spaceData.Length == 0 {
				disk.Remove(space)
			}

			file.Value = File{Id: fileData.Id, Length: fileData.Length, IsFile: false}
			break
		}
		fileId--
	}
	return checksumList(disk)
}

func checksumList(disk *list.List) (sum int) {
	idx := 0
	for cursor, cursorFile := elementAndFile(disk.Front()); cursor != nil; cursor, cursorFile = elementAndFile(cursor.Next()) {
		if !cursorFile.IsFile {
			idx += cursorFile.Length
			continue
		}
		for i := 0; i < cursorFile.Length; i, idx = i+1, idx+1 {
			sum += idx * cursorFile.Id
		}
	}
	return
}

func elementAndFile(el *list.Element) (*list.Element, File) {
	if el == nil {
		return nil, File{}
	}
	return el, el.Value.(File)
}
