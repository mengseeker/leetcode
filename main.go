package main

func main() {

}

type Dataset struct {
	DatasetType int
	Name string
	Dir string
}

type DatasetVer struct {
	Ver int
	DataCount int
	FileCount int
	FileSize int
}

// store int elasticsearch
type Data struct {
	DataType int
	Files    []string
	Metadata interface{}
}

type Label struct {
	DataSetVerID int
	TaskType     int
	TaskID       int

	LabelValType int
	Label        string
}
