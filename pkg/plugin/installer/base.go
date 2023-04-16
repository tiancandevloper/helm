/*
Copyright The Helm Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package installer // import "github.com/tiancandevloper/helm/pkg/plugin/installer"

import (
	"path/filepath"

	"github.com/tiancandevloper/helm/pkg/helmpath"
)

type base struct {
	// Source is the reference to a plugin
	Source string
}

func newBase(source string) base {
	return base{source}
}

// Path is where the plugin will be installed.
func (b *base) Path() string {
	if b.Source == "" {
		return ""
	}
	return helmpath.DataPath("plugins", filepath.Base(b.Source))
}
