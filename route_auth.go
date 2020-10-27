package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/data"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		generateHTML(w, r, nil, "layout", "users/login")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		user, err := data.FindUserByEmail(email)
		if err != nil {
			fmt.Println("メールアドレスが間違っています")
			http.Redirect(w, r, "/login", 302)
			return
		}
		if err := user.Authenticate(password); err != nil {
			fmt.Println("パスワードが間違っています")
			http.Redirect(w, r, "/login", 302)
			return
		}
		log.Println("ログインに成功しました")
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("セッションの作成に失敗しました")
			return
		}
		cookie := http.Cookie{
			Name:     "sessionId",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Println("すでにログアウトしています")
	} else {
		// DB側のセッションを削除
		session, err := data.FindSessionByUuid(cookie.Value)
		if err != nil {
			fmt.Println("セッションの取得エラー")
			return
		}
		session.Delete()
		// クライアント側のセッションを削除
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/", 302)
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		generateHTML(w, r, nil, "layout", "users/signup")
	case "POST":
		if err := r.ParseForm(); err != nil {
			return
		}
		password, err := cryptePassword(r.PostFormValue("password"))
		if err != nil {
			fmt.Println(err)
			return
		}
		user := data.User{
			Email:           r.PostFormValue("email"),
			CryptedPassword: password,
		}
		if err := user.Create(); err != nil {
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}
