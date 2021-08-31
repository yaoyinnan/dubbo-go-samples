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

package org.apache.dubbo;

import org.apache.dubbo.config.ApplicationConfig;
import org.apache.dubbo.config.ReferenceConfig;
import org.apache.dubbo.config.RegistryConfig;
import org.apache.dubbo.rpc.service.GenericService;

import java.lang.Thread;

public class Consumer {
     private static GenericService genericService;

     public static void main(String[] args) {
          initConfig();
//          try {
//               Thread.sleep(1000);
//          } catch (InterruptedException e) {
//               e.printStackTrace();
//          }
          callGetUser();
//          callQueryUser();
     }

     private static void initConfig(){
          System.out.println("\n\n\nstart to init config");
          ApplicationConfig applicationConfig = new ApplicationConfig();
          ReferenceConfig<GenericService> reference = new ReferenceConfig<GenericService>();
          applicationConfig.setName("UserProviderGer");
          reference.setApplication(applicationConfig);
          RegistryConfig registryConfig = new RegistryConfig();
          registryConfig.setAddress("zookeeper://127.0.0.1:2181");
          reference.setRegistry(registryConfig);
          reference.setGeneric(true);
          reference.setInterface("org.apache.dubbo.UserProvider");
          genericService = reference.get();
     }

     private static void callGetUser(){
          System.out.println("\n\n\ncall GetUser");
          System.out.println("start to generic invoke");
          Object[] parameterArgs = new Object[]{"A003"};
          Object result = genericService.$invoke("GetUser", null , parameterArgs);
          System.out.println("res: " + result);
     }

     private static void callQueryUser(){
          System.out.println("\n\n\ncall queryUser");
          System.out.println("start to generic invoke");
          User user = new User();
          user.setName("Patrick");
          user.setId("id");
          user.setAge(10);
          Object[] parameterArgs = new Object[]{user};
          Object result1 = genericService.$invoke("queryUser", new String[]{"com.ikurento.user.User"} , parameterArgs);
          System.out.println("res: " + result1);
     }

}
