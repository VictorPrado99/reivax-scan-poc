package scan

import (
	"fmt"
)

const (
	SqlInjectionName = "SqlInjection"
)

func init() {
	var sInjectionScanner CodeScanner
	sInjectionScanner = &SqlInjection{
		DefaultCodeScanner{
			SqlInjectionName,
		},
	}

	GetInstance().AddScan(&sInjectionScanner)
}

type SqlInjection struct {
	DefaultCodeScanner
}

func (s SqlInjection) Run() error {
	fmt.Println("Not Implemented")
	return nil
}
