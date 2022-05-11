package analysis_output

type DefaultStaticAnalysisOutput struct {
	ScanType    string
	FilePath    string
	FileLine    int
	LineContent string
}

func (d *DefaultStaticAnalysisOutput) GetScanType() string {
	return d.ScanType
}

func (d *DefaultStaticAnalysisOutput) GetFilePath() string {
	return d.FilePath
}

func (d *DefaultStaticAnalysisOutput) GetFileLine() int {
	return d.FileLine
}

func (d *DefaultStaticAnalysisOutput) GetLineContent() string {
	return d.LineContent
}
