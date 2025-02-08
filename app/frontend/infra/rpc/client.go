// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"log"
	"sync"

	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/conf"
	frontendUtils "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/utils"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth/authservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	AuthClient authservice.Client

	once sync.Once
)

// getConsulResolver 创建并返回一个 Consul 解析器
func getConsulResolver() (discovery.Resolver, error) {
	registryAddr := conf.GetConf().Hertz.RegistryAddr
	r, err := consul.NewConsulResolver(registryAddr)
	if err != nil {
		log.Printf("Failed to create Consul resolver: %v", err)
		return nil, err
	}
	return r, nil
}

func InitClient() {
	once.Do(func() {
		initUserClient()
		initAuthClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//r, err := getConsulResolver()
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initAuthClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	AuthClient, err = authservice.NewClient("auth", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
