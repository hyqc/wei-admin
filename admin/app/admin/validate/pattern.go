package validate

// 正则规则表达式
const (
	PatternPhoneRule = "^([\\s\\S]{0})$|^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$"
	PatternPhoneMsg  = "手机号格式错误"

	PatternAdminUsernameRule = "^[a-zA-Z0-9_]+$"
	PatternAdminUsernameMsg  = "账号只能是数字，字母，下划线组合"

	PatternAdminPasswordRule = "^[\\S][\\s\\S]{4,30}[\\S]$"
	PatternAdminPasswordMsg  = "密码为6-32位非空字符串"

	PatternTrimBlankStringRule = "[\\s\\S]?|^[\\S][\\s\\S]*[\\S]$"
	PatternTrimBlankStringMsg  = "不是有效字符串"

	PatternAdminApiPathRule = "^/[a-zA-z]+(/\\w{0,}){0,}"
	PatternAdminApiPathMsg  = "不是有效路由"

	PatternAdminApiKeyRule = "^\\w+::\\w+"
	PatternAdminApiKeyMsg  = "不是有效路由键名"
)
