package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(filename string) []string {
	file, err := os.Open("./inputs/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	var lineBuffer, part []byte
	var isPrefix bool

	for {
		if part, isPrefix, err = reader.ReadLine(); err != nil {
			break
		}
		lineBuffer = append(lineBuffer, part...)
		if !isPrefix {
			lines = append(lines, string(lineBuffer[:]))
			lineBuffer = make([]byte, 0)
		}
	}
	return lines
}

func ReadInputAsByteSlice(filename string) [][]byte {
	file, err := os.Open("./inputs/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]byte
	reader := bufio.NewReader(file)
	var lineBuffer, part []byte
	var isPrefix bool

	for {
		if part, isPrefix, err = reader.ReadLine(); err != nil {
			break
		}
		lineBuffer = append(lineBuffer, part...)
		if !isPrefix {
			lines = append(lines, lineBuffer)
			lineBuffer = make([]byte, 0)
		}
	}
	return lines
}

func ReadInputAndConvert[T any](filename string, convert func([]byte) (T, error)) ([]T, error) {
	file, err := os.Open("./inputs/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []T
	reader := bufio.NewReader(file)
	var lineBuffer, part []byte
	var isPrefix bool

	for {
		if part, isPrefix, err = reader.ReadLine(); err != nil {
			break
		}
		lineBuffer = append(lineBuffer, part...)
		if !isPrefix {
			converted, err := convert(lineBuffer)
			if err != nil {
				return nil, err
			}
			lines = append(lines, converted)
			lineBuffer = make([]byte, 0)
		}
	}
	return lines, nil
}
