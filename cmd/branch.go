// Copyright © 2019 Fabian Suhrau <fabian.suhrau@me.com>
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
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "goya branch [PRO-1235]",
	Long:  `format a ticket number to a branch name for easier branch creation`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var ticket string
		if len(args) == 0 {
			return fmt.Errorf("ticket missing")
		}
		ticket = args[0]

		jiraClientURL := viper.GetString("jiraurl")
		username := viper.GetString("username")
		token := viper.GetString("password")
		if len(jiraClientURL) == 0 {
			return fmt.Errorf("goya not configured")
		}

		tp := jira.BasicAuthTransport{
			Username: username,
			Password: token,
		}
		jiraClient, err := jira.NewClient(tp.Client(), jiraClientURL) // "https://issues.apache.org/jira/"
		if err != nil {
			return err
		}
		issue, _, err := jiraClient.Issue.Get(ticket, nil)
		if err != nil {
			return err
		}

		issueType := strings.ToLower(issue.Fields.Type.Name)
		issueTypes := viper.GetStringMapString("types")
		for k, v := range issueTypes {
			issueType = strings.Replace(issueType, k, v, -1)
		}

		summary := strings.Replace(strings.ToLower(issue.Fields.Summary), " ", "-", -1)
		summary = Replace(summary, "][.!?/\\")
		branchFormat := fmt.Sprintf("%s%s-%s", issueType, ticket, summary)
		os.Stdout.WriteString(branchFormat)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
