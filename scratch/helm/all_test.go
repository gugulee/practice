package helm

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_decodeHelm3(t *testing.T) {
	data, err := ioutil.ReadFile("/tmp/1")
	if nil != err {
		t.Fatalf("%s", err)
	}

	a, err := decodeHelm3(string(data))
	if nil != err {
		t.Fatalf("%s", err)
	}

	fmt.Printf("%#v\n", a.Config)
	ioutil.WriteFile("/tmp/1", []byte(a.Manifest), 0644)
}

func Test_encodeIndex(t *testing.T) {
	data, err := ioutil.ReadFile("/tmp/index.yaml")
	if nil != err {
		t.Fatalf("%s", err)
	}

	out, err := encodeIndex(data)
	if nil != err {
		t.Fatalf("%s", err)
	}

	fmt.Println(string(out))
}

func Test_decodeIndex(t *testing.T) {
	data, err := ioutil.ReadFile("/tmp/1")
	if nil != err {
		t.Fatalf("%s", err)
	}

	err = decodeIndex(data)
	if nil != err {
		t.Fatalf("%s", err)
	}
}
