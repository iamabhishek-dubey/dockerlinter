package rules

import (
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL4005 Use SHELL to change the default shell
func validateDL4005(node *parser.Node) (rst []ValidateResult, err error) {
	for _, child := range node.Children {
		switch child.Value {
		case RUN:
			isLn := false
			for _, v := range strings.Fields(child.Next.Value) {
				switch v {
				case "ln":
					isLn = true
				case "/bin/sh":
					if isLn {
						rst = append(rst, ValidateResult{line: child.StartLine, addMsg: ""})
					}
				}
			}
		}
	}
	return rst, nil
}
