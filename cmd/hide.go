// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/OrcaLLC/napkin/hide"
	"github.com/spf13/cobra"
	"log"
)

// Flags
var inputFile string
var outputFile string

// hideCmd represents the hide command
var hideCmd = &cobra.Command{
	Use:   "hide",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Sprintln("Obfuscating %s", inputFile)
		if len(inputFile) < 1 {
			log.Fatal("No input file specified.")
		}
		if len(outputFile) < 1 {
			log.Fatal("No out file specified.")
		}

		err := hide.HideFile(inputFile, outputFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	hideCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file to be napkin'd")
	hideCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output napkin to write on")

	RootCmd.AddCommand(hideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
