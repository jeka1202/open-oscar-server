package webapi

import (
	"net/http"
	"os"
)

const (
	// webClientDirEnv optionally points the WebAPI server at a directory of
	// static web client files. When unset, the repository-local client directory
	// is used so local development works without extra configuration.
	webClientDirEnv = "WEBCLIENT_DIR"

	defaultWebClientDir = "server/webapi/webclient"
)

func webClientDir() string {
	if dir := os.Getenv(webClientDirEnv); dir != "" {
		return dir
	}
	if _, err := os.Stat(defaultWebClientDir); err == nil {
		return defaultWebClientDir
	}
	return "webclient"
}

func webClientFileSystem() http.FileSystem {
	return http.Dir(webClientDir())
}
