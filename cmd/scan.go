/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	// "bufio"

	"github.com/VictorPrado99/reivax-scan-poc/code_scanner"
	"github.com/VictorPrado99/reivax-scan-poc/util"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your file for static vulnerabilities",
	Long:  `Scan your files looking for static vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]
		scanManager := code_scanner.GetInstance()

		util.CheckDirectory(directory, true)

		// extension, err := cmd.Flags().GetStringSlice("extension")
		// println("Extension = ", err)
		// outputFormat, err := cmd.Flags().GetStringSlice("output")
		// println("output format = ", err)

		libRegEx := util.BuildRegexFilterByExtension()

		files := util.GetFiles(directory, libRegEx)

		outputManager := scanManager.RunScanners(*files)

		outputManager.GenerateOutput()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringSliceP("extension", "e", []string{"js", "html", "go"}, `Set extensions you wanna scan for. Don't use .  e.g. [-e=".js,.go"] use instead [-e="js,go"]`)
	scanCmd.Flags().StringSliceP("output", "o", []string{"json, plain"}, `Set extensions you wanna scan for. Usage [-o="json, plain"]`)

	// scanCmd.Flags()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
