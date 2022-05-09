package scan

import (
	"fmt"
)

const (
	SqlInjectionName = "SqlInjection"
)

func init() {
	var sInjectionScanner CodeScanner
	sInjectionScanner = &SqlInjection{}

	GetInstance().AddScan(&sInjectionScanner)
}

type SqlInjection struct {
}

func (s SqlInjection) Run() error {
	fmt.Println("Not Implemented")
	return nil
}

func (s SqlInjection) GetName() string {
	return SqlInjectionName
}
