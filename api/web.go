package api

import (
	"renqabs/gpb/api/helper"
	"renqabs/gpb/common"
	"renqabs/gpb/web"
	"net/http"
)

func WebStatic(w http.ResponseWriter, r *http.Request) {
	if _, ok := web.WEB_PATH_MAP[r.URL.Path]; ok || r.URL.Path == common.PROXY_WEB_PREFIX_PATH {
		http.StripPrefix(common.PROXY_WEB_PREFIX_PATH, http.FileServer(web.GetWebFS())).ServeHTTP(w, r)
	} else {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
	}
}
