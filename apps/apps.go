package apps

import "net/http"

// List of apps
var List = map[string]func(w http.ResponseWriter, r *http.Request){
	"hello": hello,
}
