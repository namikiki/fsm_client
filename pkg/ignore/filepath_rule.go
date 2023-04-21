package ignore

import (
	"path/filepath"
)

type filePathRule struct {
	expr string
}

func newFilePathRule(expr string) (Rule, error) {
	return &filePathRule{
		expr: expr,
	}, nil
}

func (r *filePathRule) Match(s string) bool {
	matched, _ := filepath.Match(r.expr, s)
	return matched
}

func (r *filePathRule) Expression() string {
	return r.expr
}
