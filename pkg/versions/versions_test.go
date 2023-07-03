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

// versions contains all the global variables for setting version
package versions

import "testing"

func TestVersionIsWhatIsSet(t *testing.T) {
	version = "my_version"
	if version != GetVersion() {
		t.Errorf("incorrect version %v instead of %v", GetVersion(), version)
	}
}

func TestGitShaIsWhatIsSet(t *testing.T) {
	gitSha = "my_version"
	if gitSha != GetGitSha() {
		t.Errorf("incorrect gitsha %v instead of %v", GetGitSha(), gitSha)
	}
}
