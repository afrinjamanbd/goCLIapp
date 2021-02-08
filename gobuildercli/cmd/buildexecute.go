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
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/kopy"
	"github.com/spf13/cobra"
)

var (
	builddir string
	copydir  string
	exe      string
	skip     string
)

// buildexecuteCmd represents the buildexecute command
var buildexecuteCmd = &cobra.Command{
	Use:   "buildexecute",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Buildexecuted successfully")

		if copydir != "" {
			startcopy(builddir, copydir)
		}
		if exe != "" {
			// makefile with exe
		}
		/*		if excludefile == true {
				skip = "ok"
			} */

	},
}

func currentpath() string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	return pwd
}

func init() {
	rootCmd.AddCommand(buildexecuteCmd)

	rootCmd.PersistentFlags().StringVar(&builddir, "builddir", "", "if the user wants to copy to a specific directory.")
	rootCmd.PersistentFlags().StringVar(&copydir, "copydir", "", "source directory to copy from")
	rootCmd.PersistentFlags().StringVar(&exe, "exe", "", "if the user wants create binary with specific name")

	buildexecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildexecuteCmd.Flags().BoolP("exclude-tests", "exclude-tests", true, "exclude all test files")

}

func startcopy(bdir, cdir string) {

	IgnoreFilesOrFolders := []string{}
	if currentpath() != bdir || bdir == "" {
		bdir = currentpath()
	}

	src, dst, msg := filepath.FromSlash(bdir), filepath.FromSlash(cdir), ""
	msg = `Starts copying the entire directory or a folder: `
	fmt.Println(msg, src)
	itrlog.Infow(msg, "src", src, "log_time", time.Now().Format(itrlog.LogTimeFormat))

	filesCopied, foldersCopied, err := kopy.CopyDir(src, dst, true, IgnoreFilesOrFolders)
	if err != nil {
		fmt.Println(err)
		itrlog.Errorw("error", "err", err, "log_time", time.Now().Format(itrlog.LogTimeFormat))
		return
	}
	msg = `Successfully copied the entire directory `
	fmt.Println(msg, src, ", Number of Folders Copied: ", filesCopied, " Number of Files Copied: ", foldersCopied)
	itrlog.Infow(msg, "src", src, "dst", dst, "folder_copied", filesCopied, "files_copied", foldersCopied, "log_time", time.Now().Format(itrlog.LogTimeFormat))
}

//Skip is
func Skip(t *testing.T) {
	if skip == "ok" {
		t.Skip("Skipping testing in CI environment")
	}
}
