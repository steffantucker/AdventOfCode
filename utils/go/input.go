package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func GetInputFile(year, day int) (string, error) {
	sessionKey := os.Getenv("adventofcodecookie")
	dayString := fmt.Sprint(day)
	yearString := fmt.Sprint(year)

	fileName := fmt.Sprintf("day%02d", day)
	appData, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("AppData folder not found: %v", err)
	}
	filePath := path.Join(appData, "AoC", yearString, fileName)
	var file *os.File
	_, err = os.Stat(filePath)
	if err == nil {
		return filePath, nil
	} else if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModeDir); err != nil {
			return "", err
		}
		file, err = os.Create(filePath)
		if err != nil {
			return "", err
		}
	} else {
		return "", err
	}

	inputURL := "https://adventofcode.com/" + yearString + "/day/" + dayString + "/input"
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		return "", err
	}
	cookie := http.Cookie{
		Name:  "session",
		Value: sessionKey,
	}
	req.AddCookie(&cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func GetInput(year, day int) (string, error) {
	fileName, err := GetInputFile(year, day)
	if err != nil {
		return "", err
	}
	data, _ := os.ReadFile(fileName)
	return string(data), nil
}

func GetString(year, day int) string {
	input, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	return input
}

// GetStringList returns an array with the input
// seperated on newlines.
//
// input:
// this
// is
// input
//
// returns:
// ["this", "is", "input"]
func GetStringList(year, day int) []string {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	return strings.Split(strings.TrimSpace(file), "\n")
}

// GetParagraphs returns an array with the input
// split on double new lines.
//
// input:
// this is input paragraph one
//
// this is input paragraph two
//
// returns:
// ["this is input paragraph one", "this is input paragraph two"]
func GetParagraphs(year, day int) []string {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	paras := strings.Split(file, "\n\n")
	for i, p := range paras {
		if len(p) == 0 {
			paras = slices.Delete(paras, i, i+1)
		} else {
			paras[i] = strings.TrimSpace(p)
		}

	}
	return paras
}

func GetGroupedNumberList(year, day int) [][]int {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	out := make([][]int, 0, 500)
	group := make([]int, 0, 20)
	for _, line := range strings.Split(file, "\n") {
		if line == "" {
			out = append(out, group)
			group = make([]int, 0, 20)
		} else {
			num, _ := strconv.Atoi(line)
			group = append(group, num)
		}
	}
	return out
}

func GetNumberMatrix(year, day int) [][]int {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	out := make([][]int, 0, 500)
	for _, line := range strings.Split(file, "\n") {
		row := make([]int, 0, 50)
		for _, num := range strings.Split(line, "") {
			n, _ := strconv.Atoi(num)
			row = append(row, n)
		}
		out = append(out, row)
	}
	return out
}
