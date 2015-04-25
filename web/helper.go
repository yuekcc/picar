package web

// 设置默认值
// 如果 in 为空，则返回 d
//
func emptyStr(in string, d string) string {
	if in == "" {
		return d
	} else {
		return in
	}
}
