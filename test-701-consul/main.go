package main

import (
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"net/http"
)

func kvTest() {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	p := &api.KVPair{Key: "REDIS_MAXCLIENztS", Value: []byte("1000")}

	_, err = kv.Put(p, nil)

	if err != nil {
		panic(err)
	}

	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("KV: %v %s\n", pair.Key, pair.Value)

}

func registerServer() {

	client, _ := api.NewClient(api.DefaultConfig())
	svc, _ := connect.NewService("my-service", client)
	defer svc.Close()

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()

}

func registerServer2() {

	client, _ := api.NewClient(api.DefaultConfig())
	svc, _ := connect.NewService("my-service", client)
	defer svc.Close()

	listener, _ := tls.Listen("tcp", ":8080", svc.ServerTLSConfig())
	defer listener.Close()

	//go acceptLoop(listener)

}

func registerServer3() {

	config := api.DefaultConfig()
	config.Address = "localhost"
	client, _ := api.NewClient(config)

	registration := new(api.AgentServiceRegistration)
	registration.ID = "1234"
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"user-tomcat"}
	registration.Address = "localhost"

	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	check.Timeout = "5s"
	check.Interval = "5s"

	registration.Check = check

	client.Agent().ServiceRegister(registration)

}

func findServer() {
	client, _ := api.NewClient(api.DefaultConfig())
	service, _, _ := client.Agent().Service("1234", nil)
	fmt.Println(service)

	services, _ := client.Agent().Services()
	fmt.Println(services)

	servvices, _ := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	fmt.Println(servvices)
}

//func registerGrpc(){
//
//	consul.new
//
//}

func main() {

}
