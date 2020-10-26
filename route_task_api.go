package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/data"
)

func handleApiTasks(w http.ResponseWriter, r *http.Request) {
	// ユーザーを取得
	userID := r.URL.Query().Get("userId")
	user, err := data.FindUserById(userID)
	if err != nil {
		fmt.Println("ユーザーの取得エラー")
		return
	}
	// ユーザーのタスク一覧を取得
	tasks, err := user.Tasks()
	if err != nil {
		fmt.Println("タスクの取得エラー")
		return
	}
	// タスクをjsonにシリアライズする
	output, err := json.MarshalIndent(&tasks, "", "\t")
	if err != nil {
		fmt.Println("Marshal時のエラー")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
