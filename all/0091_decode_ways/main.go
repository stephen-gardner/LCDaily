package main

func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	if len(s) <= 1 {
		return 1
	}
	if s[0] == '1' && s[1] != '0' {
		return 1 + numDecodings(s[2:])
	}
	if s[0] == '2' && s[1] <= '6' {
		if len(s) == 2 || (len(s) > 2 && s[2] != '0') {
			return 1 + numDecodings(s[1:])
		}
	}
	return numDecodings(s[1:])
}
