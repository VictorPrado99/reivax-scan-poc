package analysis_output

import (
	"sync"
)

const (
	plainTextType = "plain"
)

type PlainTextAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *PlainTextAnalysisOutput) GenerateOutput(outputDatas []StaticAnalysisOutput, wg *sync.WaitGroup) {
	// if wg != nil {
	// 	defer wg.Done()
	// }

	// for _, outputData := range outputDatas {
	// 	println("[", outputData.GetScanType(), `] in file "`, outputData.GetFilePath(), `" on line `, outputData.GetFileLine())
	// }
}

func init() {
	plainOutput := PlainTextAnalysisOutput{
		DefaultOutputFormat{
			plainTextType,
		},
	}

	AddOutputType(&plainOutput)
}
