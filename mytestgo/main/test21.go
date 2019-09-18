package main

import "errors"

type RegisterReq struct {
	Username       string `json:"username"`
	PasswordNew    string `json:"password_new"`
	PasswordRepeat string `json:"password_repeat"`
	Email          string `json:"email"`
}

func register(req RegisterReq) error {
	if len(req.Username) > 0 {
		if len(req.PasswordNew) > 0 && len(req.PasswordRepeat) > 0 {
			if req.PasswordNew == req.PasswordRepeat {
				if emailFormatValid(req.Email) {
					createUser()
					return nil
				} else {
					return errors.New("invalid email")
				}
			} else {
				return errors.New("password and reinput must be the same")
			}
		} else {
			return errors.New("password and password reinput must be longer than 0")
		}
	} else {
		return errors.New("length of username cannot be 0")
	}
}

func createUser()  {

}

func emailFormatValid(email string)  bool{
	return true
}
