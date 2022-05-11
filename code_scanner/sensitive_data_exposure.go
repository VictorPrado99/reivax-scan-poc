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
	scannerId := "SensDataExp"

	defaultExtensionsTypes := []string{"go", "js", "html"}

	codeScanner = &SensitiveDataExposure{
		DefaultCodeScanner{
			SensitiveDataExposureName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&codeScanner)
}

type SensitiveDataExposure struct {
	DefaultCodeScanner
}

func (s SensitiveDataExposure) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {

}
