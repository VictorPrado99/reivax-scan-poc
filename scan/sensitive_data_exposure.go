package scan

import (
	"fmt"
)

const (
	SensitiveDataExposureName = "SensitiveDataExposure"
)

func init() {
	var sensitiveScanner CodeScanner
	sensitiveScanner = &SensitiveDataExposure{}

	GetInstance().AddScan(&sensitiveScanner)
}

type SensitiveDataExposure struct {
}

func (s SensitiveDataExposure) Run() error {
	fmt.Println("Not Implemented")
	return nil
}

func (s SensitiveDataExposure) GetName() string {
	return SensitiveDataExposureName
}
