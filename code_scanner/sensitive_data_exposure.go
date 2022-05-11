package code_scanner

import (
	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
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

func (s SensitiveDataExposure) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {

}
