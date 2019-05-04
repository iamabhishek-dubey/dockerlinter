package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3005 is "Do not use apt-get upgrade or dist-upgrade."
func validateDL3005(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		if child.Value == RUN {
			isAptGet, isUpgrade := false, false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "apt-get":
					isAptGet = true
				case "upgrade":
					isUpgrade = true
				}
			}
			if isAptGet && isUpgrade {
				rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
			}
		}
	}
	return rst, nil
}
