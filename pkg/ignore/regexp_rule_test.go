package ignore

import (
	"log"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	Rex, err := regexp.Compile("~$")
	if err != nil {
		return
	}

	matchString := Rex.MatchString("/Users/zylzyl/Desktop/markdown/synctest/res/4.txt~")

	if matchString {
		log.Printf("true")
	}
}
