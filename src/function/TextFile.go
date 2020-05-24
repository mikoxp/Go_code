package function

import (
	"os"
	"fmt"
	"io/ioutil"
)

func AddLineToFile(fileName string,line string){
	f, err :=os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintln(f, line)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadTextFile(fileName string) string{
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return  string(data)
}