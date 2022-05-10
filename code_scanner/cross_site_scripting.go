package code_scanner

import (
	"io/fs"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

const (
	CrossSiteScriptingName = "Cross Site Scripting"
)

func init() {
	var crossScanner CodeScanner
	crossScanner = &CrossSiteScripting{
		DefaultCodeScanner{
			CrossSiteScriptingName,
		},
	}

	GetInstance().AddScan(&crossScanner)
}

type CrossSiteScripting struct {
	DefaultCodeScanner
}

func (c CrossSiteScripting) Run(files *[]fs.File, outputManager *analysis_output.OutputManager) {

}
