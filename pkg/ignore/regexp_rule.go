package ignore

import (
	"regexp"
)

type regexpRule struct {
	expr string
	reg  *regexp.Regexp
}

func newRegexpRule(expr string) (Rule, error) {
	reg, err := regexp.Compile(expr)
	return &regexpRule{
		expr: expr,
		reg:  reg,
	}, err
}

func (r *regexpRule) Match(s string) bool {
	return r.reg.MatchString(s)
}

func (r *regexpRule) Expression() string {
	return r.expr
}
