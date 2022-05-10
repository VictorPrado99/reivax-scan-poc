package code_scanner

import (
	"fmt"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

const (
	SensitiveDataExposureName = "SensitiveDataExposure"
)

func init() {
	var codeScanner CodeScanner
	codeScanner = &SensitiveDataExposure{
		DefaultCodeScanner{
			SensitiveDataExposureName,
		},
	}

	GetInstance().AddScan(&codeScanner)
}

type SensitiveDataExposure struct {
	DefaultCodeScanner
}

func (s SensitiveDataExposure) Run(outputManager *analysis_output.OutputManager) {
	fmt.Println("Not Implemented")
}
