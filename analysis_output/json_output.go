package analysis_output

import "sync"

const (
	jsonType = "json"
)

type JsonAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *JsonAnalysisOutput) GenerateOutput(outputData []StaticAnalysisOutput, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	println(o.GetOutputType() + " generate")
}

func init() {
	jsonOutput := JsonAnalysisOutput{
		DefaultOutputFormat{
			jsonType,
		},
	}

	AddOutputType(&jsonOutput)
}
