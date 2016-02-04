/*
 * Copyright (C) 2016 Miguel Moll
 *
 * This file is part of the Paranoid Minion
 *
 * The Paranoid Minion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Paranoid Minion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with The Paranoid Minion.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"

	"dynastic.ninja/paranoid/minion"
)

func main() {

	config, err := minion.InitConfig("minion.toml")
	if err != nil {
		fmt.Println("Unable to load config file. Error:", err)
		fmt.Println("Using minion defaults.")
	}
	minion.InitLogging(*config)

	minion.Log.Info("Minion has started. Ready Up!")
}
