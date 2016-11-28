package util

import (
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)

func ReadLines(f *os.File) (lines []string){
	defer f.Close()
	reader := bufio.NewReader(f)
	contents, _ := ioutil.ReadAll(reader)

	return strings.Split(string(contents), "\n")
}