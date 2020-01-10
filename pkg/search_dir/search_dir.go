package search_dir

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func dirBfs(dirPath string) error {

	dirPaths := []string{dirPath}

	for len(dirPaths) > 0 {
		current := dirPaths[0]
		dirPaths = dirPaths[1:]

		contents, err := ioutil.ReadDir(current)
		if err != nil {
			return err
		}

		for _, content := range contents {
			fmt.Println(joinPath(current, content.Name()))

			if content.IsDir() {
				dirPaths = append(dirPaths, joinPath(current, content.Name()))
			}
		}

	}

	return nil
}

func dirDfs(dirPath string) error {

	dirPaths := []string{dirPath}

	for len(dirPaths) > 0 {
		current := dirPaths[len(dirPaths)-1]
		dirPaths = dirPaths[:len(dirPaths)-1]

		contents, err := ioutil.ReadDir(current)
		if err != nil {
			return err
		}

		for _, content := range contents {
			fmt.Println(joinPath(current, content.Name()))

			if content.IsDir() {
				dirPaths = append(dirPaths, joinPath(current, content.Name()))
			}
		}

	}

	return nil
}

func joinPath(dirs ...string) string {
	return strings.Join(dirs, `\`)
}
