package code_scanner

import (
	"regexp"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

type SensitiveDataExposureAnalysisMethod struct {
	DefaultAnalysisMethod
	dataExposureRegex *regexp.Regexp
}

func NewSensitiveDataExposureAnalysisMethod() *SensitiveDataExposureAnalysisMethod {
	scannerId := "CrossSiteScript"
	supportedExtensionsTypes := []string{"go", "js", "java", "kt"}
	dataExposureRegex := regexp.MustCompile(`(?m)(?i) .*(Checkmarx|Hellman & Friedman| \$1.15b).*(Checkmarx|Hellman & Friedman| \$1.15b).*(Checkmarx|Hellman & Friedman| \$1.15b)`)

	return &SensitiveDataExposureAnalysisMethod{
		DefaultAnalysisMethod{
			scannerId,
			supportedExtensionsTypes,
		},
		dataExposureRegex,
	}
}

func init() {
	AddAnalysisMethod(NewCrossSiteScriptingAnalysisMethod())
}

func (c *SensitiveDataExposureAnalysisMethod) Analyse(fileContent []string, path string, scannerName string) []analysis_output.StaticAnalysisOutput {
	var listAnalysisFile []analysis_output.StaticAnalysisOutput

	println("analysing file for sensitive data exposure", path)

	for lineCounter, line := range fileContent {
		if c.dataExposureRegex.MatchString(line) {
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
