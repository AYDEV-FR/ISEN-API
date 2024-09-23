package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   persistence.CacheStore
	expire  time.Duration
	key     string
}

var _ gin.ResponseWriter = &cachedWriter{}

func cachePage(store persistence.CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		key := CreateKeyFromContext(c)
		if err := store.Get(key, &cache); err != nil {
			if !errors.Is(err, persistence.ErrCacheMiss) {
				log.Println(err.Error())
			}
			// replace writer
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)

			// Drop caches of aborted contexts
			if c.IsAborted() {
				err := store.Delete(key)
				if err != nil {
					log.Println(err.Error())
				}
			}
		} else {
			fmt.Printf("EXIST\n")
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			_, err := c.Writer.Write(cache.Data)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

// CreateKey creates a package specific key for a given string
func CreateKeyFromContext(c *gin.Context) string {
	path := url.QueryEscape(c.Request.URL.RequestURI())
	token := c.GetHeader("Token")
	return fmt.Sprintf("%s-%s", token, path)
}

func newCachedWriter(store persistence.CacheStore, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		store := w.store

		//cache responses with a status code < 300
		if w.Status() < 300 {
			val := responseCache{
				w.Status(),
				w.Header(),
				data,
			}
			_ = store.Set(w.key, val, w.expire)
		}
	}
	return ret, err
}
