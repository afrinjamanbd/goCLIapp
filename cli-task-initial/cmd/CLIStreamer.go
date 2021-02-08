/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

// CliStreamerRecord is
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

// CliRunnerRecord is
type CliRunnerRecord struct {
	// How many streamer will run.
	Run         string `csv:"Run"`
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

// CliStreamerRecord s
func (cliRunnerRecord CliRunnerRecord) CliStreamerRecord() CliStreamerRecord {
	return CliStreamerRecord{
		Title:       cliRunnerRecord.Title,
		Message1:    cliRunnerRecord.Message1,
		Message2:    cliRunnerRecord.Message2,
		StreamDelay: cliRunnerRecord.StreamDelay,
		RunTimes:    cliRunnerRecord.RunTimes,
	}
}

//CliStreamerRecordCsv is
func (cliRunnerRecord CliRunnerRecord) CliStreamerRecordCsv() string {
	cliStreamerRecords := []CliStreamerRecord{cliRunnerRecord.CliStreamerRecord()}

	out, err := gocsv.MarshalString(cliStreamerRecords)

	if err != nil {
		panic(err)
	}
	return out
}

// Csv is
func Csv(cliRunners *[]CliRunnerRecord) string {
	out, err := gocsv.MarshalString(cliRunners)

	if err != nil {
		panic(err)
	}
	return out
}

// CLIStreamerCmd represents the CLIStreamer command
var CLIStreamerCmd = &cobra.Command{
	Use:   "CLIStreamer",
	Short: "A brief description of your command CLIStreamerCmd",
	Long:  `A longer description for CLIStreamerCmd.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(CLIStreamerCmd)
}

var cliRunners []CliRunnerRecord
var fileMutex sync.Mutex

//Getclistreammermsg1 is ...
func Getclistreammermsg1() {

	args := "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,First Streamer,First1 Message,Second Msg,1,10\n2,Second Streamer,First2 Message,Second Msg,3,10"

	gocsv.UnmarshalString(args, &cliRunners)
	fmt.Print(Csv(&cliRunners))
	fmt.Println("---------------------------------")

	outputFile, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	for i, runner := range cliRunners {
		fmt.Println(i, ":")

		for runtime := 0; runtime < runner.CliStreamerRecord().RunTimes; runtime++ {
			//fmt.Print(runner.CliStreamerRecordCsv())
			fileMutex.Lock()
			io.Copy(outputFile, strings.NewReader(runner.CliStreamerRecord().Message1+"\n"))
			defer fileMutex.Unlock()
			time.Sleep(time.Second * time.Duration(runner.CliStreamerRecord().StreamDelay))

		}
		//fmt.Println(runner.CliStreamerRecord().Message1)
	}

}

//Getclistreammerms2 is a
func Getclistreammerms2() {

	args := "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n2,First Streamer,First1 Message,Second Msg,1,10\n2,Second Streamer,First2 Message,Second Msg,3,10"

	gocsv.UnmarshalString(args, &cliRunners)
	fmt.Print(Csv(&cliRunners))
	fmt.Println("---------------------------------")

	outputFile, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for i, runner := range cliRunners {
		fmt.Println(i, ":")

		for runtime := 0; runtime < runner.CliStreamerRecord().RunTimes; runtime++ {
			//fmt.Print(runner.CliStreamerRecordCsv())
			fileMutex.Lock()
			io.Copy(outputFile, strings.NewReader(runner.CliStreamerRecord().Message2+"\n"))
			defer fileMutex.Unlock()
			time.Sleep(time.Second * time.Duration(runner.CliStreamerRecord().StreamDelay))

		}
		//fmt.Println(runner.CliStreamerRecord().Message1)
	}

}
