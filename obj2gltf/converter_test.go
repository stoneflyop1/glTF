package obj2gltf

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestObjToGlTF(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	objFile := path.Join(gopath, "src/github.com/stoneflyop1/glTF/merge/models/googlesf.obj")
	fileObj, err := os.Open(objFile)
	if err != nil {
		t.Error(err)
	}
	content, err := Convert(fileObj)
	if err != nil {
		t.Error(err)
	} else {
		DEBUG(content)
	}
}

func DEBUG(i interface{}) {
	fmt.Printf("%v\n", i)
	//fmt.Printf("%v\n", i)
}
