package code_scanner

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/VictorPrado99/reivax-scan-poc/analysis_output"
	"github.com/VictorPrado99/reivax-scan-poc/util"
)

const (
	CrossSiteScriptingName = "Cross Site Scripting"
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

func (c CrossSiteScripting) Run(files *[]util.FileWrapper, outputManager *analysis_output.OutputManager) {
	var listAnalysisOutput []analysis_output.StaticAnalysisOutput

	libRegEx := regexp.MustCompile(`(?m)(?i)alert\(\)`)

	for _, file := range *files {
		log.Println("Foreach File")
		extension := filepath.Ext(file.GetPath())
		if extension == ".js" || extension == ".ts" || extension == ".html" {
			log.Println("if extension")
			oFile, err := os.Open(file.GetFileInfo().Name())
			if err == nil {
				log.Fatal("Coudn't open file ", file.GetPath())
			}
			defer oFile.Close()
			fileScanner := bufio.NewScanner(oFile)
			for lineCounter := 1; fileScanner.Scan(); lineCounter++ {
				line := fileScanner.Text()

				log.Println("Reading line ", lineCounter, " Content: ", line)

				libRegEx.MatchString(line)

				var analysisOutput analysis_output.StaticAnalysisOutput
				analysisOutput = &analysis_output.DefaultStaticAnalysisOutput{
					c.GetName(),
					file.GetPath(),
					lineCounter,
				}

				listAnalysisOutput = append(listAnalysisOutput, analysisOutput)
			}
			if fileScanner.Err() != nil {
				log.Println("Scan Err = ", fileScanner.Err())
			}
		}
	}

	log.Println("List Analysis ", listAnalysisOutput)

	outputManager.AddAnalysedDataGroup(listAnalysisOutput)
}
