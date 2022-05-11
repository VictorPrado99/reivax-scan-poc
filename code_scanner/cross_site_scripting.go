package code_scanner

const (
	CrossSiteScriptingName = "Cross Site Scripting"
)

func init() {
	var crossScanner CodeScanner

	scannerId := "CrossSiteScript"

	defaultExtensionsTypes := []string{"js", "html"}

	crossScanner = &CrossSiteScripting{
		DefaultCodeScanner{
			CrossSiteScriptingName,
			defaultExtensionsTypes,
			scannerId,
		},
	}

	GetInstance().AddScan(&crossScanner)
}

type CrossSiteScripting struct {
	DefaultCodeScanner
}
