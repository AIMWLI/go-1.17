package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogResponseBody(c *gin.Context) string {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	return w.body.String()
	//fmt.Println("Response body: " + w.body.String())
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "privationel",
	})
}
