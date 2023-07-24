package hashing

import "flag"

type InputReader interface {
	Read() (*InputVars, error)
}

type InputVars struct {
	MaxConcurrent int
	Urls          []string
}

type CliReader struct{}

const ParallelDefault = 10

func (ia *CliReader) Read() (*InputVars, error) {
	iv := new(InputVars)
	flag.IntVar(&iv.MaxConcurrent, "parallel", ParallelDefault, "Your number of parallel processes")
	flag.Parse()

	iv.Urls = flag.Args()
	return iv, nil
}
