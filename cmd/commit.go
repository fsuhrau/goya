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
	"os/exec"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "goya commit [PROJ-12346]",
	Long:  `format a jira ticket to a proper commit message`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var ticketEx = regexp.MustCompile(viper.GetString("ticket"))

		var ticket string
		if len(args) == 0 {
			cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
			out, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("ticket missing")
			}

			ticket = ticketEx.FindString(string(out))
			if len(ticket) == 0 {
				return fmt.Errorf("ticket missing")
			}
		} else {
			ticket = args[0]
		}

		ticket = strings.ToUpper(ticket)

		jiraClientURL := viper.GetString("url")
		username := viper.GetString("username")
		token := viper.GetString("password")
		if len(jiraClientURL) == 0 {
			return fmt.Errorf("goya not configured")
		}

		tp := jira.BasicAuthTransport{
			Username: username,
			Password: token,
		}
		jiraClient, err := jira.NewClient(tp.Client(), jiraClientURL)
		if err != nil {
			return err
		}
		issue, _, err := jiraClient.Issue.Get(ticket, nil)
		if err != nil {
			return err
		}

		commitMessage := fmt.Sprintf("%s: %s", ticket, issue.Fields.Summary)
		if viper.GetBool("clipboard") {
			clipboard.WriteAll(commitMessage)
		}
		os.Stdout.WriteString(commitMessage)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
