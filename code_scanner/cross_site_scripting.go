package code_scanner

import (
	"fmt"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

const (
	CrossSiteScriptingName = "CrossSiteScripting"
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

func (c CrossSiteScripting) Run(outputManager *analysis_output.OutputManager) {
	fmt.Println("Not Implemented")
}
