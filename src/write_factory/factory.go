package write_factory

import (
	fun "Go_code/src/function"
	"fmt"
)
type WriteObject struct {
	Place string
	Data string
	FileName string
}
type PlaceWriter interface {
	writeData()
}

type writeInterface interface {
	writeData()
}
type fileWrite struct {
	data string
	fileName string
}
type outWrite struct {
	data string
}
func (el fileWrite) writeData()  {
	fmt.Printf("LOG: file\n")
	fun.AddLineToFile(el.fileName,el.data)
}
func (el outWrite) writeData()  {
	fmt.Printf("LOG: out\n")
	fmt.Println(el.data)
}
func WriteTextData(w writeInterface) {
	w.writeData()
}

func WriteTextToPlace(obj WriteObject)  {
	var place=obj.Place
	switch place {
	case "FILE":
		WriteTextData(fileWrite{data:obj.Data,fileName:obj.FileName})
	case "OUT":
		WriteTextData(outWrite{data:obj.Data})
	}
}