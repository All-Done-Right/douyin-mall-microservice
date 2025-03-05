package main

import (
	"context"
	"douyin-mall/app/order/kitex_gen/cart"
	"douyin-mall/app/order/kitex_gen/order"
	"douyin-mall/app/order/kitex_gen/order/orderservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func NewOrderClient() (orderservice.Client, error) {
	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:2379
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	return orderservice.NewClient("douyin.mall.order", client.WithResolver(r)) // 指定 Resolver
}
func main() {
	c, err := NewOrderClient()
	if err != nil {
		log.Fatal(err)
	}
	req := new(order.PlaceOrderReq)
	req.UserId = 1
	req.UserCurrency = "USD"
	req.Email = "111"
	req.Address = &order.Address{
		StreetAddress: "123",
		City:          "123",
		State:         "123",
		Country:       "123",
		ZipCode:       123,
	}
	req.OrderItems = []*order.OrderItem{
		{
			Cost: 11.2,
			Item: &cart.CartItem{
				ProductId: 1,
				Quantity:  2,
			},
		},
	}
	fmt.Println(len(req.OrderItems))

	resp, err := c.PlaceOrder(context.Background(), req)
	resp2, err := c.ListOrder(context.Background(), &order.ListOrderReq{
		UserId: 1,
	})
	fmt.Println(111)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp2)
	fmt.Println(111)
	fmt.Println(resp)
}
