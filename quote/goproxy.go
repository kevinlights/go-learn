package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main3() {
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct",
		"GOPRIVATE=git.example.com",
	)
	g.ProxiedSUMDBs = []string{"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org"}
	http.ListenAndServe("localhost:8080", g)
}
