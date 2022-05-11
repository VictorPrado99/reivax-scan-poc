package code_scanner

import (
	"sync"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
)

type CodeScanner interface {
	GetName() string
	Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager)
	GetDefaultExtensionsTypes() []string
	GetScannerId() string
}

type DefaultCodeScanner struct {
	ScannerName            string
	DefaultExtensionsTypes []string
	ScannerId              string
}

func (dcs *DefaultCodeScanner) GetName() string {
	return dcs.ScannerName
}

func (dcs *DefaultCodeScanner) GetScannerId() string {
	return dcs.ScannerId
}

func (dcs *DefaultCodeScanner) GetDefaultExtensionsTypes() []string {
	return dcs.DefaultExtensionsTypes
}

func (c *DefaultCodeScanner) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {
	var listAnalysisOutput []analysis_output.StaticAnalysisOutput

	for _, file := range *files {
		for _, analyseMethod := range GetAnalysisMethods(c.ScannerId, file.GetExtension()) {
			if analyseMethod != nil {
				listAnalysisOutput = append(listAnalysisOutput, analyseMethod.Analyse(file.GetFileContent(), file.GetPath(), c.ScannerName)...)
			}
		}
	}

	outputManager.AddAnalysedDataGroup(listAnalysisOutput)
}

type ScanManager struct {
	scanners           []CodeScanner
	scannersDictionary map[string]CodeScanner
}

func (manager *ScanManager) AddScan(scan *CodeScanner) {
	manager.scanners = append(manager.scanners, *scan)
}

func (manager *ScanManager) GetScanners() []CodeScanner {
	return manager.scanners
}

func (manager *ScanManager) GetScanner(scannerId string) CodeScanner {
	return manager.scannersDictionary[scannerId]
}

func (manager *ScanManager) RunScanners(files *[]util.FileWrapper) *analysis_output.OutputManager {
	outputManager := analysis_output.OutputManager{}

	wg := sync.WaitGroup{}
	for _, codeScanner := range manager.GetScanners() {
		println("Running", "scanner", codeScanner.GetScannerId())
		codeScanner.Run(files, &outputManager)
		// wg.Add(1)
		// go func(codeScanner CodeScanner) {
		// 	fmt.Println("Running " + codeScanner.GetName())
		// 	codeScanner.Run(files, &outputManager)
		// 	defer wg.Done()
		// }(codeScanner)
	}
	wg.Wait()
	return &outputManager
}

var singletonInstance *ScanManager

var once = sync.Once{}

func GetInstance() *ScanManager {
	if singletonInstance == nil {
		once.Do(func() {
			singletonInstance = &ScanManager{
				scanners: []CodeScanner{},
			}
		},
		)
	}

	return singletonInstance

}
