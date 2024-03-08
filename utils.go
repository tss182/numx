package numx

func getBase(baseType string) string {
	switch baseType {
	case typeStr:
		return baseStr
	case typeNumber:
		return baseNumberOnly
	}
	return ""
}
