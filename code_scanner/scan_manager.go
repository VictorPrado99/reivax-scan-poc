package code_scanner

import (
	"fmt"
	"io/fs"
	"sync"
)

type CodeScanner interface {
	GetName() string
	Run()
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

func (manager *ScanManager) RunScanners(files *[]fs.FileInfo) {

	wg := sync.WaitGroup{}
	for _, codeScanner := range manager.GetScanners() {
		wg.Add(1)
		go func(codeScanner CodeScanner) {
			defer wg.Done()
			fmt.Println("Running " + codeScanner.GetName())
			codeScanner.Run()
		}(codeScanner)
	}
	wg.Wait()
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
