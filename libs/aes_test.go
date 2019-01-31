package libs

import "testing"

func TestDecrypt(t *testing.T) {
	result := "U2FsdGVkX19qHlY+dJbME4FYyuNEpxLd0bBlDVBxLwU="
	x := Decrypt(result,"123")
	if x!="123" {
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	text := "123"
	x := Encrypt(text,"123")
	m := Decrypt(x,"123")
	t.Log(x)
	if m!=text {
		t.Fail()
	}
}