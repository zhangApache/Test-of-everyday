package main

import (
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type RegisterReq3 struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username       string   `validate:"gt=0"`
	// 同上
	PasswordNew    string   `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string   `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email          string   `validate:"email"`
}
//%USERPROFILE%\go

func validate(req RegisterReq3) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		//doSomething()
		return err
	}
	return nil
}

func main()  {
	var req = RegisterReq3 {
		Username       : "Xargin",
		PasswordNew    : "ohno",
		PasswordRepeat : "ohn",
		Email          : "alex@abc.com",
	}

	err := validate(req)
	fmt.Println(err)
}


