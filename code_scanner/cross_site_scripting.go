package code_scanner

import (
	"log"
	"path/filepath"

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

func (c CrossSiteScripting) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {
	var listAnalysisOutput []analysis_output.StaticAnalysisOutput

	for _, file := range *files {
		for _, analyseMethod := range GetAnalysisMethods(c.ScannerId, filepath.Ext(file.GetPath())) {
			listAnalysisOutput = append(listAnalysisOutput, analyseMethod.Analyse(file.GetFileContent(), file.GetPath(), c.ScannerName)...)
		}
	}

	log.Println("List Analysis ", listAnalysisOutput)

	outputManager.AddAnalysedDataGroup(listAnalysisOutput)
}
