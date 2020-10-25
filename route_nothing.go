package main

import (
	"net/http"
)

func doNothing(w http.ResponseWriter, r *http.Request) {
	// faviconのリクエスト時は何もしない
}
