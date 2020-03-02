// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package cloudfoundry

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "cloudfoundry", asset.ModuleFieldsPri, AssetCloudfoundry); err != nil {
		panic(err)
	}
}

// AssetCloudfoundry returns asset data.
// This is the base64 encoded gzipped contents of module/cloudfoundry.
func AssetCloudfoundry() string {
	return "eJzclbFy2zAQRHt+xY178wNYpIgzmXGXwqljCFhKGIEAAxwk8+8zoESbpCDFCWUVVqHiAL5d7uGIe9qiq0gaF1XtolW+K4hYs0FFd+PyXUGkEKTXLWtnK/pSEBE9pC30/bCHGqeiQUHkYSACKsJLC68bWBamIKo1jApV/+g9WdHgRDz9uGtR0dq72B4rGekpbUwUbftay8HOAg+/pw0Swmgp0jLxBtSAvZakA4kQnNSCoWiveVOOHp37GXvSalIebG3R7Z2fr10wNxh8/Eau7q2NrJYnUUhnWWgLvySQ51fK8wAMb+RjNoFq75rpgRiHc+5IvCs8G1hYiV/aKrxkgzTOrv8pxceEGjIcBMa9XiExA7Ers6ZkG8tWctZNbZyYr/ylqQ8/flIMYg1q4WVKZ428cIPG+a5cdYxwnSzoa2KlMGKAGgQuif+OjsXHWBA7oY1YGVz0oXTYfmAEgZ0/24Be+yYJnNh4m+toefFU94zJTPeVG0x0+r/uBzERh3E+vseZ7sGwuFLLkrDSdQ2P9PVYgffA4cYwIjCxbjB2RNjBMjkpo/dQeYPspjkuNdjzaCdMBNXO5xMaxPtti45VT/iEh6p/r3zHotX5e+C/hRPxHcLzbtGSG+hp0JoJ/wkAAP//aoqyCQ=="
}