// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ansible

import (
	"path/filepath"

	"github.com/Medium/operator-sdk/internal/scaffold"
	"github.com/Medium/operator-sdk/internal/scaffold/input"
)

const RolesFilesDir = "files" + filePathSep + ".placeholder"

type RolesFiles struct {
	StaticInput
	Resource scaffold.Resource
}

// GetInput - gets the input
func (r *RolesFiles) GetInput() (input.Input, error) {
	if r.Path == "" {
		r.Path = filepath.Join(RolesDir, r.Resource.LowerKind, RolesFilesDir)
	}
	r.TemplateBody = rolesFilesDirPlaceholder
	return r.Input, nil
}

const rolesFilesDirPlaceholder = ``
