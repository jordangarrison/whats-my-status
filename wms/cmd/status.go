/*
Copyright © 2021 Jordan Garrison <jordan.andrew.garrison@gmail.com>

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
	"strings"

	"github.com/jordangarrison/whats-my-status/status"
	util "github.com/jordangarrison/whats-my-status/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config util.Config
	// statusCmd represents the status command
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Set your status",
		Long: `You can set your status using this subcommand. Examples:

	$ wms status -e ":smile:" "I'm happy!"
	$ wms status -e :car: -t 30m Running an errand

You can also use aliases which are predefined in your config file (see https://github.com/jordangarrison/whats-my-status for more info).

Clear is a built in

	$ wms status clear

You can also define your aliases in your config file and run them of course.

	$ wms status focus

You can run wms without the status subcommand as well to perform a status update, status is the default command.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("No status provided")
				os.Exit(1)
			}

			// Get the status
			viper.Set("status.StatusMessage", strings.Join(args, " "))

			// Create config struct
			err := viper.Unmarshal(&config)
			cobra.CheckErr(err)

			// Add clear status alias to config
			config.Aliases = config.GetStatusAliases()

			// Check aliases
			for _, alias := range config.Aliases {
				if alias.Name == args[0] {
					config.Status.StatusMessage = alias.Status.StatusMessage
					config.Status.Time = alias.Status.Time
					config.Status.Emoji = alias.Status.Emoji
					break
				}
			}

			// Get the time
			if config.Status.Time != "" {
				epoch, err := util.GetEpochTime(config.Status.Time)
				cobra.CheckErr(err)
				config.Status.Epoch = epoch
				iso8601, err := util.GetISO8601Time(config.Status.Time)
				cobra.CheckErr(err)
				config.Status.ISO8601 = iso8601
			}

			fmt.Printf("Status Message: %+v\nEmoji: %+v\nTime: %+v\n", config.Status.StatusMessage, config.Status.Emoji, config.Status.Epoch)

			// Set status
			err = status.SetStatus(config)
			cobra.CheckErr(err)
		},
	}
)

func init() {
	rootCmd.AddCommand(statusCmd)
	// Emoji flag
	statusCmd.PersistentFlags().StringP("emoji", "e", "", "Emoji to use")
	viper.BindPFlag("status.emoji", statusCmd.PersistentFlags().Lookup("emoji"))

	// Time flag
	statusCmd.PersistentFlags().StringP("time", "t", "", "Time to use")
	viper.BindPFlag("status.time", statusCmd.PersistentFlags().Lookup("time"))
}
