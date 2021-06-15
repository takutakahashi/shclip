package operation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/takutakahashi/snip/pkg/parse"
)

func Write(data map[string]parse.File) error {
	for filepath, buf := range data {
		if filepath == "stdout" {
			fmt.Println(string(buf.Data))
			continue
		}
		if err := os.MkdirAll(path.Dir(filepath), 0775); err != nil {
			return err
		}
		if err := ioutil.WriteFile(filepath, buf.Data, buf.Perm); err != nil {
			return err
		}
	}
	return nil
}

func PrintList(out []ListOutput) error {
	for _, o := range out {
		fmt.Println(o.Name)
	}
	return nil
}
