package sec

import (
	"log"
	"testing"
)

func TestT1(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	if err := SaveJWT("zzz"); err != nil {
		log.Println(err)
		return
	}

	jwt, err := ReadJWT()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(jwt)

}
