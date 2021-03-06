// Copyright 2018 The cpchain authors
// This file is part of cpchain.
//
// cpchain is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cpchain is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cpchain. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"bitbucket.org/cpchain/chain/commons/log"
	"github.com/naoina/toml"
	"github.com/urfave/cli"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      dumpConfig,
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   " ",
		Description: `The dumpconfig command shows configuration values.`,
	}
)

func dumpConfig(ctx *cli.Context) error {
	cfg, _ := newConfigNode(ctx)
	err := toml.NewEncoder(os.Stdout).Encode(cfg)
	if err != nil {
		log.Fatalf("Encoding config to TOML failed: %v", err)
	}
	return nil
}
