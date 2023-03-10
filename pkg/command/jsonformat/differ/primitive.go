package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/pulumi/terraform/pkg/command/jsonformat/computed"

	"github.com/pulumi/terraform/pkg/command/jsonformat/computed/renderers"
)

func (change Change) computeAttributeDiffAsPrimitive(ctype cty.Type) computed.Diff {
	return change.asDiff(renderers.Primitive(change.Before, change.After, ctype))
}
