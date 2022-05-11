package code_scanner

import (
	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
)

const (
	CrossSiteScriptingName = "Cross Site Scripting"
)

func init() {
	var crossScanner CodeScanner

	scannerId := "CrossSiteScript"

	defaultExtensionsTypes := []string{"js", "html"}

	crossScanner = &CrossSiteScripting{
		DefaultCodeScanner{
			CrossSiteScriptingName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&crossScanner)
}

type CrossSiteScripting struct {
	DefaultCodeScanner
}

func (c *CrossSiteScripting) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {
	var listAnalysisOutput []analysis_output.StaticAnalysisOutput

	for _, file := range *files {
		println("cross_site_scripting.go ", "reading", file.GetPath())
		for _, analyseMethod := range GetAnalysisMethods(c.ScannerId, file.GetExtension()) {
			if analyseMethod != nil {
				listAnalysisOutput = append(listAnalysisOutput, analyseMethod.Analyse(file.GetFileContent(), file.GetPath(), c.ScannerName)...)
			}
		}
	}

	outputManager.AddAnalysedDataGroup(listAnalysisOutput)
}
