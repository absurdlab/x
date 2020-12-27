package httpx

import (
	"net/http"
	"net/url"
	"strings"
)

// URLQuery returns a URL query by its name and already trimmed for space.
func URLQuery(r *http.Request, name string) string {
	return strings.TrimSpace(r.URL.Query().Get(name))
}

// URLValues provides a clean way to construct url.Values.
func URLValues(kv ...string) url.Values {
	if len(kv)%2 != 0 {
		panic("odd parameters not permitted")
	}

	u := url.Values{}
	for i := 0; i < len(kv); i += 2 {
		u.Add(kv[i], kv[i+1])
	}

	return u
}
