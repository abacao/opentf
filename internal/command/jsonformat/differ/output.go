// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/placeholderplaceholderplaceholder/opentf/internal/command/jsonformat/computed"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/command/jsonformat/computed/renderers"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/command/jsonformat/structured"
)

func ComputeDiffForOutput(change structured.Change) computed.Diff {
	if sensitive, ok := checkForSensitiveType(change, cty.DynamicPseudoType); ok {
		return sensitive
	}

	if unknown, ok := checkForUnknownType(change, cty.DynamicPseudoType); ok {
		return unknown
	}

	jsonOpts := renderers.RendererJsonOpts()
	return jsonOpts.Transform(change)
}
