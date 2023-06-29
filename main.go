//	Copyright 2023 Dremio Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// main is the entry point for the statcapn cli
package main

import (
	"errors"
	"flag"
	"log"

	"github.com/rsvihladremio/statcapn/pkg"
)

func main() {
	args := ArgParse()
	if err := pkg.SystemMetrics(args); err != nil {
		log.Fatalf("unable to collect %v", errors.Unwrap(err))
	}
}

func ArgParse() pkg.Args {
	var intervalSeconds int
	var durationSeconds int
	var outFile string

	flag.IntVar(&intervalSeconds, "i", 1, "number of seconds between execution of collection")
	flag.IntVar(&durationSeconds, "d", 60, "number of seconds for duration of all collection")
	flag.Parse()
	if flag.NArg() > 0 {
		outFile = flag.Arg(0)
	}
	return pkg.Args{
		IntervalSeconds: intervalSeconds,
		DurationSeconds: durationSeconds,
		OutFile:         outFile,
	}
}
