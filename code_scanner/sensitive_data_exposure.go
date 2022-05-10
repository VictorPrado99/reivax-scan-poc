package code_scanner

import (
	"io/fs"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

const (
	SensitiveDataExposureName = "Sensitive Data Exposure"
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

func (s SensitiveDataExposure) Run(files *[]fs.File, outputManager *analysis_output.OutputManager) {

}
