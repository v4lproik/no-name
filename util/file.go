package util

import (
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)

func ReadLines(path string) ([]string, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, nil
	}

	defer f.Close()
	reader := bufio.NewReader(f)
	contents, _ := ioutil.ReadAll(reader)

	return strings.Split(string(contents), "\n"), nil
}