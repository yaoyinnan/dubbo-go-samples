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
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/protocol/dubbo"
	hessian "github.com/apache/dubbo-go-hessian2"
	gxlog "github.com/dubbogo/gost/log"
)

import (
	"github.com/apache/dubbo-go-samples/api"
)

var grpcGreeterImpl = new(api.GreeterClientImpl)

func init() {
	config.SetConsumerService(grpcGreeterImpl)

	referenceConfig := config.NewReferenceConfig(
		config.WithReferenceInterface("com.apache.dubbo.sample.basic.IGreeter"),
		config.WithReferenceProtocolName("dubbo"),
		config.WithReferenceRegistry("demoZk"),
	)

	consumerConfig := config.NewConsumerConfig(
		config.WithConsumerReferenceConfig("greeterImpl", referenceConfig),
	)

	registryConfig := config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")

	rootConfig := config.NewRootConfig(
		config.WithRootRegistryConfig("zkRegistryKey", registryConfig),
		config.WithRootConsumerConfig(consumerConfig),
	)

	if err := rootConfig.Init(); err != nil {
		panic(err)
	}
}

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/helloworld/go-client/conf/dubbogo.yml
func main() {
	config.Load()
	time.Sleep(3 * time.Second)

	logger.Info("start to test dubbo")
	req := &api.HelloRequest{
		Name: "laurence",
	}
	reply, err := grpcGreeterImpl.SayHello(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}

	logger.Infof("client response result: %v\n", reply)

	gxlog.CInfo("\n\ncall getUser")
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
	gxlog.CInfo("\n\n\nstart to generic invoke")
	resp, err := grpcGreeterImpl.GetRPCService().(*config.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser",
			[]string{"java.lang.String"},
			[]hessian.Object{"A003"},
		},
	)
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("res: %+v\n", resp)
	gxlog.CInfo("success!")

}
