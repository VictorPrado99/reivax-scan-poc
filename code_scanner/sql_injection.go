package code_scanner

const (
	SqlInjectionName = "Sql Injection"
)

func init() {
	var sInjectionScanner CodeScanner

	scannerId := "SqlInjection"

	defaultExtensionsTypes := []string{"go", "js", "html"}

	sInjectionScanner = &SqlInjection{
		DefaultCodeScanner{
			SqlInjectionName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&sInjectionScanner)
}

type SqlInjection struct {
	DefaultCodeScanner
}
