package ignore

import (
	"log"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	matchString, err := regexp.MatchString("(\\.(swp|swo)$)|(^\\.DS_Store$)|(~$)", "/Users/zylzyl/go/src/fsm_client/pkg/mock/test/jkh~")
	if err != nil {
		return
	}
	if matchString {
		log.Printf("true")
	}
}
