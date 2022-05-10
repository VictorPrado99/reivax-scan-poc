package analysis_output

const (
	plainTextType = "plain"
)

type PlainTextAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *PlainTextAnalysisOutput) GenerateOutput() {
	println(o.GetOutputType() + " generate")
}

func init() {
	plainOutput := PlainTextAnalysisOutput{
		DefaultOutputFormat{
			plainTextType,
		},
	}

	AddOutputType(&plainOutput)
}
