package analysis_output

type DefaultStaticAnalysisOutput struct {
	scanType string
	filePath string
	fileLine int
}

func (d *DefaultStaticAnalysisOutput) GetScanType() string {
	return d.scanType
}

func (d *DefaultStaticAnalysisOutput) GetFilePath() string {
	return d.filePath
}

func (d *DefaultStaticAnalysisOutput) GetFileLine() int {
	return d.fileLine
}
