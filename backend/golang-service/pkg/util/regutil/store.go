package regutil

import (
	"context"
	"net/http"
	"sync"
)

func InitCtxStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := new(sync.Map)
		r = r.WithContext(context.WithValue(r.Context(), storeKey, store))

		next.ServeHTTP(w, r)
	})
}

func Set(r *http.Request, key string, value interface{}) {
	r.Context().Value(storeKey).(*sync.Map).Store(key, value)
}

func Get(r *http.Request, key string) interface{} {
	if val, ok := r.Context().Value(storeKey).(*sync.Map).Load(key); ok {
		return val
	} else {
		return nil
	}
}
