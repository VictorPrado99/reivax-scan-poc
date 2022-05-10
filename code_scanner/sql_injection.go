package code_scanner

import (
	"io/fs"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

const (
	SqlInjectionName = "Sql Injection"
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

func (s SqlInjection) Run(files *[]fs.File, outputManager *analysis_output.OutputManager) {

}
