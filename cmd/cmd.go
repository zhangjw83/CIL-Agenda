package cmd

import (
	"entity"
	"fmt"
	"log"
	"os"
	"strings"
)
   //判断用户的登录状态错误信息 
func printWrongLoginState(action string, required bool) int {
	var s string
	if required {
		s = "login"
	} else {
		s = "logout"
	}
	util.PrintfErr("Action %s requires an %s state\n", action, s)
	return int(err.WrongLoginState)
}

   //用户登录 
func loadLogin(us entity.Users) *entity.User {
	u, e := model.LoadLogin(us)
	if e == err.OK {
		return u
	}
	log.Fatalf("something wrong with login file, error: %d", int(e))
	return nil
}

  //用户注册
func Register(user, pass, mail, phone string) int {
	users := model.LoadUsers()
	passhash := util.PrettyHash(pass)
	log.Printf("password hash for '%s': %s\n", pass, passhash)
	if !users.Add(&entity.User{
		Username: user,
		Password: passhash,
		Mail:     mail,
		Phone:    phone,
	}) {
		util.PrintfErr("there's another user with username %s\n", user)
		return int(err.DuplicateUser)
	}
	model.StoreUser(users)
	return 0
}

// 登陆的命令
func Login(user, pass string) int {
	users := model.LoadUsers()
	if loadLogin(users) != nil {
		return printWrongLoginState("login", false)
	}
	if model.Login(users, user, pass) != err.OK {
		util.PrintlnErr("Authentication Fail")
		return int(err.AuthenticateFail)
	}
	return 0
}

//退出的命令
func Logout() int {
	if model.Logout() {
		fmt.Printf("logout success\n")
	} else {
		return printWrongLoginState("logout", true)
	}
	return 0
}
