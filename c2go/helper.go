package c2go

import (
	"fmt"
	"strings"
)

func IsWordChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_'
}

func HasWord(s, word string) bool {
	for {
		idx := strings.Index(s, word)
		if idx < 0 {
			return false
		}
		beforeOk := idx == 0 || !IsWordChar(rune(s[idx-1]))
		afterOk := idx+len(word) >= len(s) || !IsWordChar(rune(s[idx+len(word)]))
		if beforeOk && afterOk {
			return true
		}
		s = s[idx+1:]
	}
}

func ReplaceWord(s, old, new string) string {
	var result strings.Builder
	i := 0
	for i < len(s) {
		idx := strings.Index(s[i:], old)
		if idx < 0 {
			result.WriteString(s[i:])
			break
		}
		pos := i + idx
		beforeOk := pos == 0 || !IsWordChar(rune(s[pos-1]))
		afterPos := pos + len(old)
		afterOk := afterPos >= len(s) || !IsWordChar(rune(s[afterPos]))
		if beforeOk && afterOk {
			result.WriteString(s[i:pos])
			result.WriteString(new)
			i = afterPos
		} else {
			result.WriteString(s[i : pos+1])
			i = pos + 1
		}
	}
	return result.String()
}

func ExtractIdentifiers(s string) []string {
	var result []string
	var buf strings.Builder
	for _, ch := range s {
		if IsWordChar(ch) {
			buf.WriteRune(ch)
		} else if buf.Len() > 0 {
			result = append(result, buf.String())
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		result = append(result, buf.String())
	}
	return result
}

func ContainsDigit(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}

func IsULiteral(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'u' || s[i] == 'U' {
			if i+1 < len(s) && (s[i+1] == '"' || s[i+1] == '\'') {
				return true
			}
		}
	}
	return false
}

func IsLargeHex(s string) bool {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(strings.ToLower(s), "0x") {
		return false
	}
	hexPart := s[2:]
	count := 0
	for _, ch := range hexPart {
		if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') {
			count++
		} else {
			break
		}
	}
	return count >= 9
}

func IsCStyleCast(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 || s[0] != '(' {
		return false
	}
	depth := 0
	typeEnd := -1
	for i, ch := range s {
		if ch == '(' {
			depth++
		} else if ch == ')' {
			depth--
			if depth == 0 {
				typeEnd = i
				break
			}
		}
	}
	if typeEnd < 0 || typeEnd == 1 {
		return false
	}
	typeName := s[1:typeEnd]
	if len(typeName) == 0 {
		return false
	}
	if typeName[0] < 'A' || typeName[0] > 'Z' {
		return false
	}
	for _, ch := range typeName[1:] {
		if !IsWordChar(ch) {
			return false
		}
	}
	rest := strings.TrimSpace(s[typeEnd+1:])
	return len(rest) > 0 && rest[0] == '('
}

func StripCStyleCast(s string) string {
	s = strings.TrimSpace(s)
	if !IsCStyleCast(s) {
		return s
	}
	depth := 0
	typeEnd := -1
	for i, ch := range s {
		if ch == '(' {
			depth++
		} else if ch == ')' {
			depth--
			if depth == 0 {
				typeEnd = i
				break
			}
		}
	}
	if typeEnd < 0 {
		return s
	}
	return strings.TrimSpace(s[typeEnd+1:])
}

func IsCharLiteral(s string) bool {
	if len(s) < 3 {
		return false
	}
	return s[0] == '\'' && s[len(s)-1] == '\''
}

func StripNumberSuffix(s string) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return s
	}
	lastAlphaIdx := n - 1
	for lastAlphaIdx >= 0 {
		ch := runes[lastAlphaIdx]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			lastAlphaIdx--
		} else {
			break
		}
	}
	if lastAlphaIdx < n-1 && lastAlphaIdx >= 0 {
		prev := runes[lastAlphaIdx]
		if (prev >= '0' && prev <= '9') || (prev >= 'a' && prev <= 'f') || (prev >= 'A' && prev <= 'F') {
			return string(runes[:lastAlphaIdx+1])
		}
	}
	return s
}

func ParseSizeofType(s string) (typeName string, found bool) {
	prefix := "sizeof("
	if !strings.Contains(s, prefix) {
		return "", false
	}
	start := strings.Index(s, prefix)
	if start < 0 {
		return "", false
	}
	parenStart := start + len(prefix)
	if parenStart >= len(s) {
		return "", false
	}
	ch := s[parenStart]
	if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_') {
		return "", false
	}
	depth := 1
	end := -1
	for i := parenStart; i < len(s); i++ {
		if s[i] == '(' {
			depth++
		} else if s[i] == ')' {
			depth--
			if depth == 0 {
				end = i
				break
			}
		}
	}
	if end < 0 {
		return "", false
	}
	typeName = strings.TrimSpace(s[parenStart:end])
	if len(typeName) == 0 {
		return "", false
	}
	for _, ch := range typeName {
		if !IsWordChar(ch) {
			return "", false
		}
	}
	return typeName, true
}

func HasNotParen(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '!' {
			j := i + 1
			for j < len(s) && s[j] == ' ' {
				j++
			}
			if j < len(s) && s[j] == '(' {
				return true
			}
		}
	}
	return false
}

func HasDoubleNotParen(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == '!' && s[i+1] == '!' {
			j := i + 2
			for j < len(s) && s[j] == ' ' {
				j++
			}
			if j < len(s) && s[j] == '(' {
				return true
			}
		}
	}
	return false
}

func FindLargeNum(s string) (numStr string, found bool) {
	runes := []rune(s)
	i := 0
	for i < len(runes) {
		if runes[i] >= '0' && runes[i] <= '9' {
			start := i
			for i < len(runes) && runes[i] >= '0' && runes[i] <= '9' {
				i++
			}
			beforeOk := start == 0 || !IsWordChar(runes[start-1])
			afterOk := i >= len(runes) || !IsWordChar(runes[i])
			if beforeOk && afterOk && (i-start) >= 10 {
				return string(runes[start:i]), true
			}
		} else {
			i++
		}
	}
	return "", false
}

func FindShiftAmount(s string) (amount int, found bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == '<' && s[i+1] == '<' {
			j := i + 2
			for j < len(s) && s[j] == ' ' {
				j++
			}
			numStart := j
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			if j > numStart {
				var val int
				fmt.Sscanf(s[numStart:j], "%d", &val)
				return val, true
			}
		}
	}
	return 0, false
}

func needsUint64(body string) bool {
	lower := strings.ToLower(body)
	for i := 0; i < len(lower)-1; i++ {
		if lower[i] == '0' && i+1 < len(lower) && lower[i+1] == 'x' {
			j := i + 2
			hexStart := j
			for j < len(lower) && ((lower[j] >= '0' && lower[j] <= '9') || (lower[j] >= 'a' && lower[j] <= 'f')) {
				j++
			}
			if j > hexStart {
				hexStr := lower[hexStart:j]
				if len(hexStr) > 8 {
					return true
				}
				if len(hexStr) == 8 && hexStr[0] >= '8' {
					return true
				}
			}
		}
	}
	if _, found := FindLargeNum(body); found {
		return true
	}
	return false
}
