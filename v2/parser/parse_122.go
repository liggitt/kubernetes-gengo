//go:build go1.22
// +build go1.22

/*
Copyright 2024 The Kubernetes Authors.

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

package parser

import (
	gotypes "go/types"

	"k8s.io/gengo/v2/types"
)

func (p *Parser) walkAliasType(u types.Universe, in gotypes.Type) *types.Type {
	switch t := in.(type) {
	case *gotypes.Alias:
		// handle alias types specifically (only in go1.22, default in go1.23+)
		name := goNameToName(t.String())
		out := u.Type(name)
		if out.Kind != types.Unknown {
			return out
		}
		rhs := p.walkType(u, nil, t.Rhs())
		if rhs.Kind == types.Alias || rhs.Kind == types.Interface {
			// avoid nesting aliases, or aliases to interfaces
			out = rhs
		} else {
			out.Kind = types.Alias
			out.Underlying = rhs
		}
		return out
	default:
		return nil
	}
}
