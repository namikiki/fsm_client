package ignore

import (
	"fmt"
	"path/filepath"
)

type filePathRule struct {
	expr string
}

func newFilePathRule(expr string) (Rule, error) {
	if _, err := filepath.Match(expr, ""); err != nil {
		return nil, fmt.Errorf("parse %s rule failed, expression=%s, %w", "filePathSwitch", expr, err)
	}
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
