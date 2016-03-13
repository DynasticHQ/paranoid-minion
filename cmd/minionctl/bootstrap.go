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
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bootstrapCmd sets the minion up the initial communication with the server.
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Boostrap the minion.",
	Long:  `Bootstrap the minion with its initial server communication.`,
	RunE:  Bootstrap,
}

// bootstrap flags
var (
	apiHostValue   string
	validatorValue string
)

const (
	ApiHost        = "server.apihost"
	ApiHostFlag    = "apihost"
	ApiHostDefault = "localhost:5555"

	Validator        = "minion.validator"
	ValidatorFlag    = "Validator"
	ValidatorDefault = "validator.pem"
)

func init() {

	bootstrapCmd.PersistentFlags().StringVarP(&apiHostValue, ApiHostFlag, "a", ApiHostDefault, "API the minion will connect to.")
	bootstrapCmd.PersistentFlags().StringVarP(&validatorValue, ValidatorFlag, "v", ValidatorDefault, "Validator key used to authenticate with the server.")

	viper.BindPFlag(ApiHost, bootstrapCmd.PersistentFlags().Lookup(ApiHostFlag))
	viper.BindPFlag(Validator, bootstrapCmd.PersistentFlags().Lookup(ValidatorFlag))
}

func Bootstrap(cmd *cobra.Command, args []string) error {

	InitializeConfig()

	logrus.Info("Bootstrapping the minion.")
	logrus.Info(viper.GetString(ApiHost))
	logrus.Info(viper.GetString(Validator))

	return nil
}
