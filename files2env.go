/**
 * Author: Robert Jensen
 * File: secrets.go
 */

package files2env

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Import(path string) {
	// walk through the directory and read each file
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// read the file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// create a env variable with the file name and content
			varName := info.Name()
			fmt.Printf("var %s = %q\n", varName, content)
			os.Setenv(string(varName), string(content))
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
