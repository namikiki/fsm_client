package ignore

import (
	"fsm_client/pkg/types"
)

type Ignore struct {
	rules []Rule
}

func NewIgnore(ignoreConfig *types.Ignore) (*Ignore, error) {
	rules, err := parseIgnoreConfig(ignoreConfig)
	return &Ignore{rules}, err
}

func (ig *Ignore) Match(s string) bool {
	for _, rule := range ig.rules {
		if rule.Match(s) {
			return true
		}
	}
	return false
}

//
//func (ig *Ignore) AddFilepathRule(expr string) error {
//	rule, err := newFilePathRule(expr)
//	if err != nil {
//		return err
//	}
//	ig.rules = append(ig.rules, rule)
//	return nil
//}
//
//func (ig *Ignore) AddRegexpRule(expr string) error {
//	rule, err := newRegexpRule(expr)
//	if err != nil {
//		return err
//	}
//	ig.rules = append(ig.rules, rule)
//	return nil
//}

func parseIgnoreConfig(ignoreConfig *types.Ignore) ([]Rule, error) {
	var rs []Rule

	for _, ig := range ignoreConfig.Filepath {
		rule, err := newFilePathRule(ig)
		if err != nil {
			return nil, err
		}
		rs = append(rs, rule)
	}

	for _, ig := range ignoreConfig.Regexp {
		rule, err := newRegexpRule(ig)
		if err != nil {
			return nil, err
		}
		rs = append(rs, rule)
	}

	return rs, nil
}

type Rule interface {
	Match(s string) bool
	Expression() string
}
