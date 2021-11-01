package util

import (
	"strings"
	"unicode"
)

var (
	commonInitialisms                  = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	commonInitialismsUpperCaseReplacer = toUpperCaseReplacer(commonInitialisms)
	commonInitialismsLowerCaseReplacer = toLowerCaseReplacer(commonInitialisms)
	_                                  = commonInitialismsLowerCaseReplacer
	_                                  = commonInitialismsUpperCaseReplacer
)

var uppercaseAcronym = map[string]string{
	"ID": "id",
}

func toUpperCaseReplacer(sms []string) *strings.Replacer {
	var list []string
	for _, initialism := range sms {
		list = append(list, strings.ToLower(initialism), initialism)
	}
	replacer := strings.NewReplacer(list...)
	return replacer
}

func toLowerCaseReplacer(sms []string) *strings.Replacer {
	var list []string
	for _, initialism := range sms {
		list = append(list, initialism, strings.ToLower(initialism))
	}
	replacer := strings.NewReplacer(list...)
	return replacer
}

// Camel2Case defined TODO
func Camel2Case(name string) string {
	buffer := NewBuffer()

	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Case2Camel defined TODO
func Case2Camel(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := true
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

// ToTitle defined TODO
func ToTitle(t string) string {
	if rt := commonInitialismsUpperCaseReplacer.Replace(t); rt != t {
		return rt
	}
	return strings.Title(Case2Camel(t))
}
