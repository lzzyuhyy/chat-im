package pkg

import (
	"errors"
	"regexp"
	"unicode"
)

// 全数字返回true
func IsNumeric(str string) bool {
	for _, v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// 全大写返回true
func IsUpperCase(str string) bool {
	for _, v := range str {
		if !unicode.IsUpper(v) {
			return false
		}
	}
	return true
}

// 全小写返回true
func IsLowerCase(str string) bool {
	for _, v := range str {
		if !unicode.IsLower(v) {
			return false
		}
	}
	return true
}

// 密码合法性校验
func VerifyPwd(pwd string) error {

	pwdReg := `^[a-zA-Z][0-9a-zA-Z]{7,19}$`
	// 先判断密码是否合规
	compile, _ := regexp.Compile(pwdReg)
	if !compile.Match([]byte(pwd)) {
		return errors.New("密码不合法, 必须为8-20为大小写字母加数字组合, 首位不能为数字")
	}

	if IsNumeric(pwd) {
		return errors.New("密码不能为纯数字")
	}
	if IsUpperCase(pwd) {
		return errors.New("密码不能为纯大写字母")
	}
	if IsLowerCase(pwd) {
		return errors.New("密码不能为纯小写字母")
	}

	return nil
}

// 密码强度判断
func PwdStrongVerify(pwd string) int {
	firstLevel := `^[0-9a-zA-Z]{8,9}$`
	secondLevel := `^[0-9a-zA-Z]{10,15}$`
	thirdLevel := `^[0-9a-zA-Z]{16,20}$`
	// 包含数字，大小写字母且长度在8位到10位之间为一级
	compile, _ := regexp.Compile(firstLevel)
	if compile.Match([]byte(pwd)) {
		return 1
	}
	// 包含数字，大小写字母且长度在13位到15位之间为二级
	compile, _ = regexp.Compile(secondLevel)
	if compile.Match([]byte(pwd)) {
		return 2
	}
	// 包含数字，大小写字母且长度在15位到20位之间为三级
	compile, _ = regexp.Compile(thirdLevel)
	if compile.Match([]byte(pwd)) {
		return 3
	}

	return 0
}
