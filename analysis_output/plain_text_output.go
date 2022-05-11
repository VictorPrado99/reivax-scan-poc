package analysis_output

import (
	"strconv"

	"github.com/VictorPrado99/reivax-scan-poc/util"
)

const (
	plainTextType = "plain"
)

type PlainTextAnalysisOutput struct {
	DefaultOutputFormat
}

func (o *PlainTextAnalysisOutput) GenerateOutput(outputDatas []StaticAnalysisOutput) {
	var lines []string
	for _, outputData := range outputDatas {
		line := "[" + outputData.GetScanType() + `] in file "` + outputData.GetFilePath() + `" on line ` + strconv.Itoa(outputData.GetFileLine())
		lines = append(lines, line)
	}

	util.WriteFile(lines, "Statis Analysis.txt")
}

func init() {
	plainOutput := PlainTextAnalysisOutput{
		DefaultOutputFormat{
			plainTextType,
		},
	}

	AddOutputType(&plainOutput)
}
