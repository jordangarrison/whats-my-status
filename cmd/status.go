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
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("No status provided")
				os.Exit(1)
			}

			// Get the status
			viper.Set("status.status", strings.Join(args, " "))

			// Create config struct
			err := viper.Unmarshal(&config)
			if err != nil {
				panic(err)
			}

			// Create clear status alias
			clearStatusAlias := util.Alias{
				Name:   "clear",
				Status: util.Status{},
			}
			// Add clear status alias to config
			config.Aliases = append(config.Aliases, clearStatusAlias)

			// Check aliases
			for _, alias := range config.Aliases {
				if alias.Name == args[0] {
					config.Status.StatusMessage = alias.Status.StatusMessage
					config.Status.Time = alias.Status.Time
					config.Status.Emoji = alias.Status.Emoji
					break
				}
			}

			// Get the epoch time
			if config.Status.Time != "" {
				epoch, err := util.GetEpochTime(config.Status.Time)
				if err != nil {
					panic(err)
				}
				config.Status.Epoch = epoch
			}

			fmt.Printf("Status Message: %+v\nEmoji: %+v\nTime: %+v\n", config.Status.StatusMessage, config.Status.Emoji, config.Status.Epoch)

			// Set status
			err = status.SetStatus(config)
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Emoji flag
	statusCmd.PersistentFlags().StringP("emoji", "e", "", "Emoji to use")
	viper.BindPFlag("status.emoji", statusCmd.PersistentFlags().Lookup("emoji"))

	// Time flag
	statusCmd.PersistentFlags().StringP("time", "t", "", "Time to use")
	viper.BindPFlag("status.time", statusCmd.PersistentFlags().Lookup("time"))

	// Status argument

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
