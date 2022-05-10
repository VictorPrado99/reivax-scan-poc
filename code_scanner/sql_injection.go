package code_scanner

import (
	"fmt"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
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

func (s SqlInjection) Run(outputManager *analysis_output.OutputManager) {
	fmt.Println("Not Implemented")
}
