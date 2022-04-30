package util

import "github.com/mozillazg/go-pinyin"

func PinYinArgs() *pinyin.Args {
	// 设置获取拼音的时候保留字母
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		if r == '(' {
			return nil
		}
		if r == ')' {
			return nil
		}
		if r == '（' {
			return nil
		}
		if r == '）' {
			return nil
		}
		if r == ' ' {
			return nil
		}
		if r == ',' {
			return nil
		}
		if r == '，' {
			return nil
		}
		if r == '，' {
			return nil
		}
		if r == '【' {
			return nil
		}
		if r == '】' {
			return nil
		}
		if r == '[' {
			return nil
		}
		if r == ']' {
			return nil
		}
		if r == '{' {
			return nil
		}
		if r == '}' {
			return nil
		}
		if r == '"' {
			return nil
		}
		if r == '「' {
			return nil
		}
		if r == '」' {
			return nil
		}
		if r == '\\' {
			return nil
		}
		if r == '、' {
			return nil
		}
		if r == '。' {
			return nil
		}
		if r == '/' {
			return nil
		}
		if r == '？' {
			return nil
		}
		if r == '《' {
			return nil
		}
		if r == '》' {
			return nil
		}
		if r == '<' {
			return nil
		}
		if r == '>' {
			return nil
		}
		if r == '.' {
			return nil
		}
		if r == '`' {
			return nil
		}
		if r == '~' {
			return nil
		}
		if r == '·' {
			return nil
		}
		if r == '！' {
			return nil
		}
		if r == '!' {
			return nil
		}
		if r == '@' {
			return nil
		}
		if r == '#' {
			return nil
		}
		if r == '¥' {
			return nil
		}
		if r == '$' {
			return nil
		}
		if r == '%' {
			return nil
		}
		if r == '^' {
			return nil
		}
		if r == '&' {
			return nil
		}
		if r == '*' {
			return nil
		}
		if r == '-' {
			return nil
		}
		if r == '=' {
			return nil
		}
		if r == '+' {
			return nil
		}
		if r == '—' {
			return nil
		}
		if r == '…' {
			return nil
		}
		if r == 'α' {
			return []string{"alpha"}
		}
		if r == 'β' {
			return []string{"beta"}
		}
		if r == 'γ' {
			return []string{"gamma"}
		}
		if r == 'δ' {
			return []string{"delta"}
		}
		if r == 'ε' {
			return []string{"epsilon"}
		}
		if r == 'ζ' {
			return []string{"zeta"}
		}
		if r == 'η' {
			return []string{"eta"}
		}
		if r == 'θ' {
			return []string{"theta"}
		}
		if r == 'ι' {
			return []string{"iota"}
		}
		if r == 'κ' {
			return []string{"kappa"}
		}
		if r == 'λ' {
			return []string{"lambda"}
		}
		if r == 'μ' {
			return []string{"mu"}
		}
		if r == 'ν' {
			return []string{"nu"}
		}
		if r == 'ξ' {
			return []string{"xi"}
		}
		if r == 'ο' {
			return []string{"omicron"}
		}
		if r == 'π' {
			return []string{"pi"}
		}
		if r == 'ρ' {
			return []string{"rho"}
		}
		if r == 'σ' {
			return []string{"sigma"}
		}
		if r == 'ς' {
			return []string{"sigma"}
		}
		if r == 'τ' {
			return []string{"tau"}
		}
		if r == 'υ' {
			return []string{"upsilon"}
		}
		if r == 'φ' {
			return []string{"phi"}
		}
		if r == 'χ' {
			return []string{"chi"}
		}
		if r == 'ψ' {
			return []string{"psi"}
		}
		if r == 'ω' {
			return []string{"omega"}
		}
		return []string{string(r)}
	}
	return &a
}
