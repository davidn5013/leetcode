package grind75

// Initializing tools, tools/ds and boa into package grind75

import (
	t "github.com/davidn5013/leetcode/tools"
	"github.com/davidn5013/leetcode/tools/ds"
	boa "github.com/davidn5013/snickerboa"
)

// This just to force the import to stay after fmt
var (
	_ = ds.NewPos()
	_ = t.MaxInt(1, 2)
	_ = boa.SortString("ba")
)
