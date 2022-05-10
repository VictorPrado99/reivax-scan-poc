package code_scanner

import (
	"fmt"
	"io/fs"
	"sync"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
)

type CodeScanner interface {
	GetName() string
	Run(files *[]fs.File, outputManager *analysis_output.OutputManager)
}

type DefaultCodeScanner struct {
	ScannerName string
}

func (dcs *DefaultCodeScanner) GetName() string {
	return dcs.ScannerName
}

type ScanManager struct {
	scanners []CodeScanner
}

func (manager *ScanManager) AddScan(scan *CodeScanner) {
	manager.scanners = append(manager.scanners, *scan)
}

func (manager *ScanManager) GetScanners() []CodeScanner {
	return manager.scanners
}

func (manager *ScanManager) RunScanners(files *[]fs.File) *analysis_output.OutputManager {
	outputManager := analysis_output.OutputManager{}

	wg := sync.WaitGroup{}
	for _, codeScanner := range manager.GetScanners() {
		wg.Add(1)
		go func(codeScanner CodeScanner) {
			defer wg.Done()
			fmt.Println("Running " + codeScanner.GetName())
			codeScanner.Run(files, &outputManager)
		}(codeScanner)
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
