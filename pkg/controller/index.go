package controller

import (
	"fmt"
	"net/http"
)

func (c *ctrl) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)
	if r.URL.Path != "/" {
		c.view.NotFound(w, r)
		return
	}
	c.view.Index(w, r)
}
