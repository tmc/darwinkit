// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol to identify call graph nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchingnode?language=objc
type PFunctionStitchingNode interface {
}

// ensure impl type implements protocol interface
var _ PFunctionStitchingNode = (*FunctionStitchingNodeObject)(nil)

// A concrete type for the [PFunctionStitchingNode] protocol.
type FunctionStitchingNodeObject struct {
	objc.Object
}
