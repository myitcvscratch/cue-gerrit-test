// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package html

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("html", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Escape",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = Escape(s)
			}
		},
	}, {
		Name: "Unescape",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = Unescape(s)
			}
		},
	}},
}
