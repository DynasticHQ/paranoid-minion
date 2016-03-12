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
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	AppName       = "Paranoid Minion"
	BinaryVersion = "0.0.1-alpha" // http://semver.org/
)

func VersionString(name, version string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", name, version, runtime.Version())
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the minion's version",
	Long:  `All software has versions. This is the minion's.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(VersionString(cmd.Parent().Name(), BinaryVersion))
		return nil
	},
}
