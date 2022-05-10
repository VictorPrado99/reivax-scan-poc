package analysis_output

import (
	"sync"
)

const (
	DefaultOutputFormatType = "plain"
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

var outputsTypes map[string]OutputFormatInterface

func AddOutputType(outputType OutputFormatInterface) {
	if outputsTypes == nil {
		outputsTypes = make(map[string]OutputFormatInterface)
	}

	outputsTypes[outputType.GetOutputType()] = outputType
}

func GetOutputType(outputType string) OutputFormatInterface {
	return outputsTypes[outputType]
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

func (o *OutputManager) GenerateOutput(outputTypes ...string) {
	for _, outputType := range outputTypes {
		GetOutputType(outputType).GenerateOutput()
	}

	if len(outputsTypes) == 0 {
		GetOutputType(DefaultOutputFormatType).GenerateOutput()
	}
}
