package cart

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetCart(t *testing.T) {
	h := server.Default()
	h.GET("/cart", GetCart)
	path := "/cart"                                           
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} 
	header := ut.Header{}                                     
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

}

func TestAddCartItem(t *testing.T) {
	h := server.Default()
	h.POST("/cart", AddCartItem)
	path := "/cart"                                           
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} 
	header := ut.Header{}                                     
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))
}
