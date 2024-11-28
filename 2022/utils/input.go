package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetInputFile(year, day int) (string, error) {
	sessionKey := os.Getenv("adventofcodecookie")
	dayString := fmt.Sprint(day)
	yearString := fmt.Sprint(year)

	fileName := "day" + dayString + "/input"
	var file *os.File
	_, err := os.Stat(fileName)
	if err == nil {
		return fileName, nil
	} else if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(fileName)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	inputURL := "https://adventofcode.com/" + yearString + "/day/" + dayString + "/input"
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("cookie", "session="+sessionKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func GetInput(year, day int) (string, error) {
	fileName, err := GetInputFile(year, day)
	if err != nil {
		return "", err
	}
	data, _ := os.ReadFile(fileName)
	return string(data), nil
}

func GetStringList(year, day int) []string {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	return strings.Split(strings.Trim(file, "\n "), "\n")
}

func GetParagraphs(year, day int) []string {
	file, err := GetInput(year, day)
	if err != nil {
		log.Fatalf("Failed getting file %#v\n", err)
	}
	return strings.Split(file, "\n\n")
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
