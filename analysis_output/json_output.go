package analysis_output

import (
	"encoding/json"
	"io/ioutil"
)

const (
	jsonType = "json"
)

type JsonAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *JsonAnalysisOutput) GenerateOutput(outputDatas []StaticAnalysisOutput) {
	file, _ := json.MarshalIndent(outputDatas, "", " ")

	_ = ioutil.WriteFile("static_analysis.json", file, 0644)

}

func init() {
	jsonOutput := JsonAnalysisOutput{
		DefaultOutputFormat{
			jsonType,
		},
	}

	AddOutputType(&jsonOutput)
}
