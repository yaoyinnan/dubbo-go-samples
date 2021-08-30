/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config/generic"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster_impl"
	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/config/generic"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	"dubbo.apache.org/dubbo-go/v3/protocol/dubbo"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	_ "dubbo.apache.org/dubbo-go/v3/registry/zookeeper"

	hessian "github.com/apache/dubbo-go-hessian2"
)

var (
	appName         = "dubbo.io"
	referenceConfig config.ReferenceConfig
)

func init() {
	registryConfig := &config.RegistryConfig{
		Protocol: "zookeeper",
		Address:  "127.0.0.1:2181",
	}

	referenceConfig = config.ReferenceConfig{
		InterfaceName: "org.apache.dubbo.UserProvider",
		Cluster:       "failover",
		Registry:      []string{"zk"},
		Protocol:      dubbo.DUBBO,
		Generic:       "protobuf-json",
	}

	rootConfig := config.NewRootConfig(config.WithRootRegistryConfig("zk", registryConfig))
	_ = rootConfig.Init()
	_ = referenceConfig.Init(rootConfig)
	referenceConfig.GenericLoad(appName)
}

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/generic/protobufjson/go-client/conf/dubbogo.yml
func main() {
	callGetUser()

	initSignal()
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP,
		syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(10*time.Second, func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("app exit now...")
			return
		}
	}
}

func callGetUser() {
	logger.Infof("\n\ncall callGetUser")
	logger.Infof("start to generic invoke")
	resp, err := referenceConfig.GetRPCService().(*generic.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser",
			[]string{},
			// "{\"id\":1}" is a JSON string that represents an UserRequestType object shown in the below:
			// &UserRequestType{
			// 		Id: 1
			// }
			[]hessian.Object{"{\"id\":1}"},
		},
	)
	if err != nil {
		panic(err)
	}
	logger.Infof("res: %+v\n", resp)
	logger.Infof("success!")
}
