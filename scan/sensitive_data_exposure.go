package scan

import (
	"fmt"
)

const (
	SensitiveDataExposureName = "SensitiveDataExposure"
)

func init() {
	var sensitiveScanner CodeScanner
	sensitiveScanner = &SensitiveDataExposure{
		DefaultCodeScanner{
			SensitiveDataExposureName,
		},
	}

	GetInstance().AddScan(&sensitiveScanner)
}

type SensitiveDataExposure struct {
	DefaultCodeScanner
}

func (s SensitiveDataExposure) Run() error {
	fmt.Println("Not Implemented")
	return nil
}

func (s SensitiveDataExposure) GetName() string {
	return SensitiveDataExposureName
}
