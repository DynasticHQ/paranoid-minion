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

package minion

import "github.com/BurntSushi/toml"

// MinionConfig describes the config for the minion.
type MinionConfig struct {
	Log LogConfig
}

// Log describes the log group in a TOML format
type LogConfig struct {
	Level  string
	Type   string
	Format string
}

// MinionConfig constatns
const (
	LogLevelInfo    = "info"
	LogLevelDebug   = "debug"
	LogLevelWarning = "warning"
	LogLevelError   = "error"

	LogFormatJSON = "json"
	LogFormatText = "text"

	LogTypeStdout = "stdout"
	LogTypeFile   = "file"
)

var Config MinionConfig

// InitConfig initalizes the MinionConfig by reading the path config.
// Returns an error if something went wrong.
func InitConfig(path string) error {

	Config := NewConfig()
	err := LoadConfigFile(path, &Config)
	if err != nil {
		return err
	}

	return nil
}

// LoadConfig takes a string path to load a TOML file and saves the data in mc.
// It returns an error if something went wrong.
func LoadConfigFile(path string, mc *MinionConfig) error {

	_, err := toml.DecodeFile(path, mc)
	if err != nil {
		return err
	}

	return nil
}

// NewConfig returns a MinionConfig with config defaults for the minion.
func NewConfig() MinionConfig {

	mc := MinionConfig{}

	mc.Log.Level = LogLevelInfo
	mc.Log.Type = LogTypeFile
	mc.Log.Format = LogFormatText

	return mc
}
