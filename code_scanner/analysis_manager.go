package code_scanner

import (
	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
)

func GetAnalysisMethods(codeScannerId string, fileExtensions ...string) []AnalysisMethod {
	var analysisMethodsList []AnalysisMethod

	for _, fileExtension := range fileExtensions {
		key := BuildKey(codeScannerId, fileExtension)
		analysisMethod := getAnalysisMethod(key)
		if !util.ContainTypeInList(analysisMethod, &analysisMethodsList) {
			analysisMethodsList = append(analysisMethodsList, analysisMethod)
		}
	}

	if len(fileExtensions) == 0 {
		for _, fileExtension := range GetInstance().GetScanner(codeScannerId).GetDefaultExtensionsTypes() {
			key := BuildKey(codeScannerId, fileExtension)
			analysisMethod := getAnalysisMethod(key)
			if !util.ContainTypeInList(analysisMethod, &analysisMethodsList) {
				analysisMethodsList = append(analysisMethodsList, analysisMethod)
			}
		}
	}

	return analysisMethodsList
}

func BuildKey(codeScannerId string, fileExtension string) string {
	return codeScannerId + "|" + fileExtension
}

type AnalysisMethod interface {
	GetAnalysisKeys() []string
	Analyse(fileContent []string, path string, scannerName string) []analysis_output.StaticAnalysisOutput
}

type DefaultAnalysisMethod struct {
	CodeScannerId      string
	ExtensionSupported []string
}

func (c *DefaultAnalysisMethod) GetAnalysisKeys() []string {
	var keys []string
	for _, extension := range c.ExtensionSupported {
		keys = append(keys, BuildKey(c.CodeScannerId, extension))
	}
	return keys
}

var analysisMethodDictionary map[string]AnalysisMethod

func AddAnalysisMethod(analysisMethod AnalysisMethod) {
	if analysisMethodDictionary == nil {
		analysisMethodDictionary = make(map[string]AnalysisMethod)
	}

	for _, key := range analysisMethod.GetAnalysisKeys() {
		analysisMethodDictionary[key] = analysisMethod
	}

}

func getAnalysisMethod(key string) AnalysisMethod {
	return analysisMethodDictionary[key]
}
