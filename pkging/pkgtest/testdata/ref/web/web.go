package web

import (
	"net/http"

	"github.com/WuErPing/pkger"
)

func Serve() {
	pkger.Stat("github.com/WuErPing/pkger:/README.md")
	dir := http.FileServer(pkger.Dir("/public"))
	http.ListenAndServe(":3000", dir)
}
