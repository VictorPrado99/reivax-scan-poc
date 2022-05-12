package analysis_output

import (
	"fmt"
	"log"
	"sync"
)

const (
	DefaultOutputFormatType = "plain"
)

type OutputFormatInterface interface {
	GenerateOutput(outputData []StaticAnalysisOutput)
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
	value := outputsTypes[outputType]
	if value == nil {
		log.Fatal("Output doesn't exist")
	}
	return outputsTypes[outputType]
}

type StaticAnalysisOutput interface {
	GetScanType() string
	GetFilePath() string
	GetFileLine() int
	GetLineContent() string
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

func (o *OutputManager) HasDataToOutput() bool {
	return len(o.analysedData) > 0
}

func (o *OutputManager) GenerateOutput(outputTypes ...string) {
	if o.HasDataToOutput() {
		wg := sync.WaitGroup{}
		for _, outputType := range outputTypes {
			wg.Add(1)
			go func(outputType string) {
				defer wg.Done()
				GetOutputType(outputType).GenerateOutput(o.analysedData)
			}(outputType)
		}

		if len(outputsTypes) == 0 {
			GetOutputType(DefaultOutputFormatType).GenerateOutput(o.analysedData)
		}

		wg.Wait()
	} else {
		fmt.Println("Wasn't possible to detect anything")
	}
}
