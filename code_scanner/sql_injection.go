package code_scanner

import (
	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
)

const (
	SqlInjectionName = "Sql Injection"
)

func init() {
	var sInjectionScanner CodeScanner

	scannerId := "SqlInjection"

	defaultExtensionsTypes := []string{"go", "js", "html"}

	sInjectionScanner = &SqlInjection{
		DefaultCodeScanner{
			SqlInjectionName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&sInjectionScanner)
}

type SqlInjection struct {
	DefaultCodeScanner
}

func (s SqlInjection) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {

}
