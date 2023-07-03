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
	"fmt"
	"log"
	"math"
	"os"

	"github.com/rsvihladremio/statcapn/pkg"
	"github.com/rsvihladremio/statcapn/pkg/versions"
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

	fs := flag.NewFlagSet("statcapn", flag.ExitOnError)

	fs.IntVar(&intervalSeconds, "i", 1, "number of seconds between execution of collection")
	fs.IntVar(&durationSeconds, "d", math.MaxInt, "number of seconds for duration of all collection")

	// Customize the usage message
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "statcapn %s-%s\n\nstandard usage:\n\tstatcapn -i <interval> -d <duration_seconds> metrics.txt\n\nFor json output:\n\tstatcapn -i <interval> -d <duration_seconds> metrics.json\n\nflags:\n\n", versions.GetVersion(), versions.GetGitSha())
		fs.PrintDefaults()
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if fs.NArg() > 0 {
		outFile = fs.Arg(0)
	}
	return pkg.Args{
		IntervalSeconds: intervalSeconds,
		DurationSeconds: durationSeconds,
		OutFile:         outFile,
	}
}
