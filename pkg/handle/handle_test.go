package handle

import (
	"fsm_client/pkg/database"
	"log"
	"testing"
)

func TestT1(t *testing.T) {

	connect := database.NewGormSQLiteConnect()

	handle := NewHandle(nil, connect, nil)
	err := handle.ScannerPathToUpload("C:\\Users\\surflabom\\Desktop\\课程实践\\结课报告", "123")
	if err != nil {
		log.Println(err)
		return
	}
	//handle.ScannerPathToUpload("C:\\Users\\surflabom\\Desktop\\MyDediServer", "123")
}
