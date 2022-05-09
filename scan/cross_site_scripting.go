package scan

import (
	"fmt"
)

const (
	CrossSiteScriptingName = "CrossSiteScripting"
)

func init() {
	var crossScanner CodeScanner
	crossScanner = &CrossSiteScripting{}

	GetInstance().AddScan(&crossScanner)
}

type CrossSiteScripting struct {
}

func (c CrossSiteScripting) Run() error {
	fmt.Println("Not Implemented")
	return nil
}

func (c CrossSiteScripting) GetName() string {
	return CrossSiteScriptingName
}
