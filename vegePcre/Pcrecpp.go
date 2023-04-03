package vegePcre

func PcreppReplaceImpl(pattern string, repl string, src string, flags string) string {
	return Pcrepp_Replace(pattern, repl, src, flags)
}
