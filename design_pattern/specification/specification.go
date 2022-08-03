package specification

import (
	"fmt"
	"strings"
)

type AclMatchSpecification interface {
	Rule() (string, error)
}

type AndSpecification struct {
	specifications []AclMatchSpecification
}

func NewAndSpecification(specifications ...AclMatchSpecification) AclMatchSpecification {
	return AndSpecification{
		specifications: specifications,
	}
}

// Rule generate acl match rule like 'ip4.src == $test.allow.as && ip4.src != $test.except.as && 12345 <= tcp.dst <= 12500 && outport == @ovn.sg.test_sg && ip'
func (s AndSpecification) Rule() (string, error) {
	var rules []string
	for _, specification := range s.specifications {
		rule, err := specification.Rule()
		if err != nil {
			return "", fmt.Errorf("generate rule %s: err", rule)
		}
		rules = append(rules, rule)
	}

	return strings.Join(rules, " && "), nil
}

type aclRuleKv struct {
	key      string
	value    string
	maxValue string
	effect   string
}

func NewKV(key, effect, value, maxValue string) AclMatchSpecification {
	return aclRuleKv{
		key:      key,
		effect:   effect,
		value:    value,
		maxValue: maxValue,
	}
}

// Rule generate acl match rule like
// 'ip4.src == $test.allow.as'
// or 'ip4.src != $test.except.as'
// or '12345 <= tcp.dst <= 12500'
// or 'tcp.dst == 13500'
// or 'outport == @ovn.sg.test_sg && ip'
func (kv aclRuleKv) Rule() (string, error) {
	// key must exist at least
	if len(kv.key) == 0 {
		return "", fmt.Errorf("acl rule key is required")
	}

	// like 'ip'
	if len(kv.effect) == 0 || len(kv.value) == 0 {
		return kv.key, nil
	}

	// like 'tcp.dst == 13500' or 'ip4.src == $test.allow.as'
	if len(kv.maxValue) == 0 {
		return fmt.Sprintf("%s %s %s", kv.key, kv.effect, kv.value), nil
	}

	// like '12345 <= tcp.dst <= 12500'
	return fmt.Sprintf("%s %s %s %s %s", kv.value, kv.effect, kv.key, kv.effect, kv.maxValue), nil
}
