package extension

// If returns trueVal if isTrue, else return falseVal
func If[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

func MultiIf[T any](cond1 bool, ret1 T, cond2 bool, ret2 T, defaultRet T) T {
	if cond1 {
		return ret1
	} else if cond2 {
		return ret2
	} else {
		return defaultRet
	}
}
