package scan

import (
	"fmt"
)

const (
	CrossSiteScriptingName = "CrossSiteScripting"
)

func init() {
	var crossScanner CodeScanner
	crossScanner = &CrossSiteScripting{
		DefaultCodeScanner{
			CrossSiteScriptingName,
		},
	}

	GetInstance().AddScan(&crossScanner)
}

type CrossSiteScripting struct {
	DefaultCodeScanner
}

func (c CrossSiteScripting) Run() error {
	fmt.Println("Not Implemented")
	return nil
}
