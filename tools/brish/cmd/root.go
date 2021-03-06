// Copyright © 2019 Yesphet
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
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "brish",
	Short: "Little Brish, help for brightFFmpeg",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(genNewCmd())
	rootCmd.AddCommand(genConfigCmd())
	rootCmd.AddCommand(genPushCmd())
	rootCmd.AddCommand(genPullCmd())
	rootCmd.AddCommand(genGatherCmd())

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Use != "config" {
			readConfig()
		}
	}

}

func CheckFatal(err error) {
	if err != nil {
		Fatalf(err.Error())
	}
}

func CheckFatalf(err error, msg string, args ...interface{}) {
	if err != nil {
		Fatalf(msg+"\n\t "+err.Error(), args...)
	}
}

func Fatalf(s string, args ...interface{}) {
	fmt.Printf("error: "+s+"\n", args...)
	os.Exit(1)
}
