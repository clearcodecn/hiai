package zipfile

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	dst, err := ZipFile("/tmp/v2ray", "abcdefg")
	fmt.Println(dst, err)
	if err != nil {
		return
	}
	//	defer os.Remove(dst)
}
