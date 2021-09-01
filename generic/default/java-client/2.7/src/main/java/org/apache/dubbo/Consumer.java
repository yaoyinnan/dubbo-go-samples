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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Consumer {
     private static GenericService genericService;

     public static void main(String[] args) {
        initConfig();
        callGetUser();
        callQueryUser();
         callQueryUsers();
//         callGetOneUser();
     }

     private static void initConfig(){
        System.out.println("\n\n\nstart to init config\n\n\n");
        ApplicationConfig applicationConfig = new ApplicationConfig();
        ReferenceConfig<GenericService> reference = new ReferenceConfig<GenericService>();
        applicationConfig.setName("user-info-server");
        reference.setApplication(applicationConfig);
        RegistryConfig registryConfig = new RegistryConfig();
        registryConfig.setAddress("zookeeper://127.0.0.1:2181");
        reference.setRegistry(registryConfig);
        reference.setGeneric(true);
        reference.setInterface("org.apache.dubbo.UserProvider");
        genericService = reference.get();
     }

     private static void callGetUser(){
        System.out.println("\n\n\nCall GetUser");
        System.out.println("Start to generic invoke");
        Object[] parameterArgs = new Object[]{"A003"};
        Object result = genericService.$invoke("GetUser", new String[]{"java.lang.String"} , parameterArgs);
          System.out.println("\n\n\n" + "res: " + result + "\n\n\n");
     }

     private static void callQueryUser(){
        System.out.println("\n\n\nCall QueryUser");
        System.out.println("Start to generic invoke");
        User user = new User();
        user.setName("Patrick");
        user.setId("id");
        user.setAge(10);
        Object[] parameterArgs = new Object[]{user};
        Object result = genericService.$invoke("QueryUser", new String[]{"org.apache.dubbo.User"} , parameterArgs);
        System.out.println("\n\n\n" + "res: " + result + "\n\n\n");
     }

     private static void callQueryUsers(){
        System.out.println("\n\n\nCall QueryUsers");
        System.out.println("Start to generic invoke");
        ArrayList<User> userArr = new ArrayList<User>();
        userArr.add(new User("A001", "Patrick", 10));
        userArr.add(new User("A002", "xavier-niu", 24));
        Object[] parameterArgs = new Object[]{userArr};
        Object result = genericService.$invoke("QueryUsers", new String[]{"java.util.ArrayList"} , parameterArgs);
        System.out.println("\n\n\n" + "res: " + result + "\n\n\n");
     }

     private static void callGetOneUser(){
        System.out.println("\n\n\nCall GetOneUser");
        System.out.println("Start to generic invoke");
        Object[] parameterArgs = new Object[]{"null"};
        Object result = genericService.$invoke("GetOneUser", new String[]{"null"} , parameterArgs);
        System.out.println("\n\n\n" + "res: " + result + "\n\n\n");
     }
}
