package code_scanner

import (
	"log"
	"regexp"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

type CrossSiteScriptingAnalysisMethod struct {
	DefaultAnalysisMethod
	libRegEx *regexp.Regexp
}

func NewCrossSiteScriptingAnalysisMethod() *CrossSiteScriptingAnalysisMethod {
	scannerId := "CrossSiteScript"
	supportedExtensionsTypes := []string{"js", "html"}
	libRegEx := regexp.MustCompile(`(?m)(?i)alert\(\)`)

	return &CrossSiteScriptingAnalysisMethod{
		DefaultAnalysisMethod{
			scannerId,
			supportedExtensionsTypes,
		},
		libRegEx,
	}
}

func init() {
	AddAnalysisMethod(NewCrossSiteScriptingAnalysisMethod())
}

func (c *CrossSiteScriptingAnalysisMethod) Analyse(fileContent []string, path string, scannerName string) []analysis_output.StaticAnalysisOutput {
	var listAnalysisFile []analysis_output.StaticAnalysisOutput

	for lineCounter, line := range fileContent {
		log.Println("Validating file ", path)
		if c.libRegEx.MatchString(line) {
			analysisOutput := &analysis_output.DefaultStaticAnalysisOutput{
				scannerName,
				path,
				(lineCounter + 1),
				line,
			}

			listAnalysisFile = append(listAnalysisFile, analysisOutput)
		}
	}
	return listAnalysisFile
}
