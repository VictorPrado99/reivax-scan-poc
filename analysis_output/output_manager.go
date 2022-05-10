package analysis_output

import (
	"sync"
)

type OutputFormatInterface interface {
	GenerateOutput()
	GetOutputType() string
}

type DefaultOutputFormat struct {
	outputType string
}

func (dof *DefaultOutputFormat) GetOutputType() string {
	return dof.outputType
}

var outputsTypes []OutputFormatInterface

func AddOutputType(outputType OutputFormatInterface) {
	outputsTypes = append(outputsTypes, outputType)
}

func GetOutputTypes() []OutputFormatInterface {
	return outputsTypes
}

type StaticAnalysisOutput interface {
	GetScanType() string
	GetFilePath() string
	GetFileLine() int
}

type OutputManager struct {
	lock         sync.Mutex
	analysedData []StaticAnalysisOutput
}

func (o *OutputManager) AddAnalysedDataGroup(analysedDataGroup []StaticAnalysisOutput) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.analysedData = append(o.analysedData, analysedDataGroup...)
}
