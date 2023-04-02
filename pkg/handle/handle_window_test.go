package handle

import "testing"

func TestT1(t *testing.T) {
	handle := NewHandle(nil, nil, nil)
	handle.ScannerPathToUpload("C:\\Users\\surflabom\\Desktop\\MyDediServer", "123")

}
