// Copyright © 2020 Fabian Suhrau <fabian.suhrau@me.com>
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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "goya pr",
	Long:  `open the create pull request page for the current command`,
	RunE: func(cmd *cobra.Command, args []string) error {
		gitCommand := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
		out, err := gitCommand.Output()
		if err != nil {
			return fmt.Errorf("not a git repository")
		}
		if len(out) == 0 {
			return fmt.Errorf("no branch detected")
		}

		bitbucketURL := viper.GetString("bitbucketurl")

		openCommand := exec.Command("open", fmt.Sprintf("%s/pull-requests?create&sourceBranch=refs/heads/%s", bitbucketURL, strings.TrimSpace(string(out))))
		openCommand.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
