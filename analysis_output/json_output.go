package analysis_output

const (
	jsonType = "json"
)

type JsonAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *JsonAnalysisOutput) GenerateOutput() {
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
