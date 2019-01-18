// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "deal date command",
	Long:  `deal date command`,
	Run: func(cmd *cobra.Command, args []string) {
		tFlag := cmd.Flags().Lookup("timestamp")
		if tFlag != nil {
			printTimestamp2Date(tFlag.Value.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dateCmd.Flags().Int64P("timestamp", "t", 0, "timestamp to date")
}

func printTimestamp2Date(value string) {
	if tm, err := strconv.ParseInt(value, 10, 64); err == nil {
		tm := time.Unix(tm, 0)
		if int(tm.Month()) < 10 {
			fmt.Printf("%d-0%d-%d %d:%d:%d\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		} else {
			fmt.Printf("%d-%d-%d %d:%d:%d\n", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		}
	}

}
