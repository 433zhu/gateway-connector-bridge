// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package main

import (
	"github.com/TheThingsNetwork/gateway-connector-bridge/cmd"
	"github.com/spf13/viper"
)

var (
	version   = "2.0.0-dev" // Corresponds with the version of the TTN backend
	gitBranch = "unknown"
	gitCommit = "unknown"
	buildDate = "unknown"
)

func main() {
	viper.Set("version", version)
	viper.Set("gitBranch", gitBranch)
	viper.Set("gitCommit", gitCommit)
	viper.Set("buildDate", buildDate)
	cmd.Execute()
}
