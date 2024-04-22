package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
	// expectPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	// if pathname != expectPathName {
	// 	t.Errorf("have %s want %s", pathname, expectPathName)
	// }

	// if pathname != expectPathName {
	// 	t.Errorf("have %s want %s", pathname, expectPathName)
	// }

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
