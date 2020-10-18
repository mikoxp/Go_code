package function

import (
	"log"
	"io/ioutil"
)

func FileInDirectory(path string)[]string{
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var result= make([]string, len(files))
	var index=0
	for _, f := range files {
		result[index]=f.Name()
		index++
	}
	return result
}
