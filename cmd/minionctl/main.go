/* Paranoid Minion
 * Copyright (C) 2016 Miguel Moll
 *
 * This software is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This software is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this software.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// minionCmd is the minions's root command.
// Every other command attached to minionCmd is a child command to it.
var minionCmd = &cobra.Command{
	Use:   "minionctl",
	Short: "Paranoid Minion Control Center",
	Long: `The interface to the Paranoid Minion.

The minion runs commands on behalf of the Paranoid Overlord.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New(fmt.Sprintf("%s is not useful directly. Please use a subcommand.\n", cmd.Use))
	},
}

// Minion flags
var (
	logLevel  string
	logFormat string
)

const (
	LogLevel        = "log.level"
	LogLevelFlag    = "loglevel"
	LogLevelDefault = "info"

	LogFormat        = "log.format"
	LogFormatFlag    = "logformat"
	LogFormatDefault = "text"
)

func init() {

	minionCmd.PersistentFlags().StringVarP(&logLevel, LogLevelFlag, "l", LogLevelDefault, "Level of verbosity used. Values: debug, info, warn, error")
	minionCmd.PersistentFlags().StringVarP(&logFormat, LogFormatFlag, "f", LogFormatDefault, "Log format. Values: json, text")

	viper.BindPFlag(LogLevel, minionCmd.PersistentFlags().Lookup(LogLevelFlag))
	viper.BindPFlag(LogFormat, minionCmd.PersistentFlags().Lookup(LogFormatFlag))

}

func main() {

	minionCmd.AddCommand(versionCmd)
	minionCmd.AddCommand(runCmd)
	minionCmd.AddCommand(bootstrapCmd)

	if err := minionCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

// InitializeConfig should be called by subcommands so that they can take advantage
// global configuration settings.
func InitializeConfig() {

	viper.SetConfigName("minion")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/paranoid/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Unable to read config file:", err, "Using default values.")
	}

	LoadDefaultSettings()

	viper.SetEnvPrefix(minionCmd.Use)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	SetupLogging()
}

func LoadDefaultSettings() {
	viper.SetDefault(LogLevel, LogLevelDefault)
	viper.SetDefault(LogFormat, LogFormatDefault)
}

func SetupLogging() {

	switch viper.GetString(LogFormat) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	logrus.SetOutput(os.Stdout) // TODO:(miguelmoll) Make a flag?

	switch viper.GetString(LogLevel) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
