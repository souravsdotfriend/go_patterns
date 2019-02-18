package functional_opts

import "testing"

func TestFunctionMethod(t *testing.T) {
	_, err := NewFile("/tmp/empty.txt")
	if err != nil {
		panic(err)
	}

	_, err = NewFile("/tmp/file.txt", UID(1000), Contents("Lorem Ipsum Dolor Amet"))
	if err != nil {
		panic(err)
	}
}