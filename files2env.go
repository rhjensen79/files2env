/**
 * Author: Robert Jensen
 * File: files2env.go
 */

package files2env

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Import(path string) {
	// walk through the directory and read each file
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Directory does not exist")
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
			log.Println(varName + " Env Variable Created")
			//fmt.Printf("var %s = %q\n", varName, content) - Can be removed
			os.Setenv(string(varName), string(content))
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
