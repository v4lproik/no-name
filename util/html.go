package util

import (
	"github.com/pmezard/go-difflib/difflib"
)

func GetDiffBetweenTwoPages(page1 string, page2 string) (float64){
	sequenceMatcher := difflib.SequenceMatcher{}
	sequenceMatcher.SetSeq1(difflib.SplitLines(page1))
	sequenceMatcher.SetSeq2(difflib.SplitLines(page2))

	return sequenceMatcher.Ratio()
}
