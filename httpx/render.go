package httpx

import (
	"encoding/json"
	"net/http"
)

// Render renders the http response.
func Render(rw http.ResponseWriter, options ...RenderOpt) {
	for _, opt := range options {
		opt(rw)
	}
}

// RenderOpt applies modification to the http.ResponseWriter
type RenderOpt func(rw http.ResponseWriter)

// SetStatus writes the given status to http.ResponseWriter
func SetStatus(status int) RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.WriteHeader(status)
	}
}

// SetApplicationJSON writes application/json to the Content-Type header.
func SetApplicationJSON() RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.Header().Set(ContentType, ApplicationJSON)
	}
}

// SetFormURLEncoded writes application/x-www-form-urlencoded to the Content-Type header.
func SetFormURLEncoded() RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.Header().Set(ContentType, ApplicationFormURLEncoded)
	}
}

// SetNoCache sets no-cache to the Cache-Control header.
func SetNoCache() RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.Header().Set(CacheControl, NoCache)
	}
}

// SetNoStore sets no-store to the Cache-Control header.
func SetNoStore() RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.Header().Set(CacheControl, NoStore)
	}
}

// SetHeader sets the given pair to http headers.
func SetHeader(header string, value string) RenderOpt {
	return func(rw http.ResponseWriter) {
		rw.Header().Set(header, value)
	}
}

// SetJSONPayload writes the given object to the body as JSON. Special cases allow
// for []byte, json.RawMessage and string to be write to the body directly without
// JSON encoding. Note this method does not set the Content-Type to application/json.
func SetJSONPayload(obj interface{}) RenderOpt {
	return func(rw http.ResponseWriter) {
		switch obj.(type) {
		case []byte:
			_, _ = rw.Write(obj.([]byte))
		case json.RawMessage:
			_, _ = rw.Write(obj.([]byte))
		case string:
			_, _ = rw.Write([]byte(obj.(string)))
		default:
			_ = json.NewEncoder(rw).Encode(obj)
		}
	}
}
