package policy

import (
	"fmt"
	"net/http"
	"path/filepath"

	httpio "github.com/intelfike/mulpage/io/http"
	htmlproc "github.com/intelfike/mulpage/policy/html"
	"github.com/intelfike/mulpage/types"
)

func Listener(w http.ResponseWriter, r *http.Request) {
	info := types.PageInfo{}
	if err := info.Init(r); err != nil {
		if err = httpio.WriteFile(w, filepath.Join("contents", info.Contents, "public", r.URL.Path)); err != nil {
			httpio.WriteFile(w, "public"+r.URL.Path)
		}
		return
	}
	// HTMLを生成して送信
	redirect, err := htmlproc.Write(w, info)
	if err != nil {
		result := "エラー:" + fmt.Sprint(err) + "\n"
		result += "エラーメッセージを記録して管理者に報告してください。"
		httpio.Write(w, result)
		return
	}
	// 必要ならリダイレクトする
	redirect.Exec(w, r)
}
