package code_scanner

const (
	SensitiveDataExposureName = "Sensitive Data Exposure"
)

func init() {
	var codeScanner CodeScanner
	scannerId := "SensDataExp"

	defaultExtensionsTypes := []string{"go", "js", "html"}

	codeScanner = &SensitiveDataExposure{
		DefaultCodeScanner{
			SensitiveDataExposureName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&codeScanner)
}

type SensitiveDataExposure struct {
	DefaultCodeScanner
}
