package scan

import (
	"sync"
)

type CodeScanner interface {
	GetName() string
	Run() error
}

type ScanManager struct {
	scanners []CodeScanner
}

func (manager *ScanManager) AddScan(scan *CodeScanner) {
	lock := &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	manager.scanners = append(manager.scanners, *scan)
}

func (Manager *ScanManager) GetScanners() []CodeScanner {
	return Manager.scanners
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
