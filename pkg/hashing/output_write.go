package hashing

import "fmt"

type OutputWriter interface {
	Write(ov OutputVars) error
}

type OutputVars struct {
	Url          string
	ResponseHash string
}

type CliWriter struct{}

func (cw *CliWriter) Write(ov OutputVars) error {
	fmt.Printf("%s\t%s\n", ov.Url, ov.ResponseHash)
	return nil
}
