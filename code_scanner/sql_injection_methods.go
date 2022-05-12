package code_scanner

import (
	"regexp"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

type SqlInjectionAnalysisMethod struct {
	DefaultAnalysisMethod
	slRegex *regexp.Regexp
}

func NewSqlInjectionAnalysisMethod() *SqlInjectionAnalysisMethod {
	scannerId := "SqlInjection"
	supportedExtensionsTypes := []string{"go"}
	slRegex := regexp.MustCompile(`(?m)(?i)("|).+?SELECT\s+?.+?\s+?FROM\s+?.+?\s+?WHERE\s+?%(s|c).?"`)

	return &SqlInjectionAnalysisMethod{
		DefaultAnalysisMethod{
			scannerId,
			supportedExtensionsTypes,
		},
		slRegex,
	}
}

func init() {
	AddAnalysisMethod(NewSqlInjectionAnalysisMethod())
}

func (c *SqlInjectionAnalysisMethod) Analyse(fileContent []string, path string, scannerName string) []analysis_output.StaticAnalysisOutput {
	var listAnalysisFile []analysis_output.StaticAnalysisOutput

	for lineCounter, line := range fileContent {
		if c.slRegex.MatchString(line) {
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
