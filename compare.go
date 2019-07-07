package riegeli

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

// ProtoDiff returns a human-readable report of the differences between two
// values, ensuring that any proto.Message values are compared correctly with
// proto.Equal.
//
// See github.com/google/go-cmp/cmp for more details.
func ProtoDiff(x, y interface{}, opts ...cmp.Option) string {
	return cmp.Diff(x, y, makeProtoOpts(opts)...)
}

func makeProtoOpts(opts []cmp.Option) []cmp.Option {
	protoOpts := append([]cmp.Option{}, opts...)
	protoOpts = append(protoOpts,
		cmp.Comparer(proto.Equal),
		ignoreProtoXXXFields,
	)
	return protoOpts
}

var ignoreProtoXXXFields = cmp.FilterPath(func(p cmp.Path) bool {
	for _, s := range p {
		if strings.HasPrefix(s.String(), ".XXX_") {
			return true
		}
	}
	return false
}, cmp.Ignore())
