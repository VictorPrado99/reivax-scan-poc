package scan

import (
	"fmt"
)

const (
	SensitiveDataExposureName = "SensitiveDataExposure"
)

func init() {
	var codeScanner CodeScanner
	codeScanner = &SensitiveDataExposure{
		DefaultCodeScanner{
			SensitiveDataExposureName,
		},
	}

	GetInstance().AddScan(&codeScanner)
}

type SensitiveDataExposure struct {
	DefaultCodeScanner
}

func (s SensitiveDataExposure) Run() error {
	fmt.Println("Not Implemented")
	return nil
}
