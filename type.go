package main

import (
	"errors"
	"fmt"
)

type Package struct {
	Module string
	Type   int
	Field  string
}

// With scope.
func (p *Package) RunValidationWarehouseCheck() bool {
	fmt.Println("run validator of wh check")
	return true
}

func (p *Package) RunValidationShipper() bool {
	fmt.Println("run validator of shipper check")
	return true
}

func (p *Package) Complete() {
	fmt.Println("completed")
}

func (p *Package) RunAfter() error {
	fmt.Println("called")
	return errors.New("error")
}
