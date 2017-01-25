package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
)

const (
	// DefaultPort is the port in which to listen if there is no PORT declared
	DefaultPort = "9000"
)

// HelloServer is the base funtion to expose text
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, World! from GO")
	if strings.Contains(strings.Join(os.Environ(), " "), "VCAP_APPLICATION") {

		appEnv, _ := cfenv.Current()

		fmt.Fprintln(w, "Super I'm running in CloudFoundry and this are my variables:")
		fmt.Fprintln(w, "ID:", appEnv.ID)
		fmt.Fprintln(w, "Index:", appEnv.Index)
		fmt.Fprintln(w, "Name:", appEnv.Name)
		fmt.Fprintln(w, "Host:", appEnv.Host)
		fmt.Fprintln(w, "Port:", appEnv.Port)
		fmt.Fprintln(w, "Version:", appEnv.Version)
		fmt.Fprintln(w, "Home:", appEnv.Home)
		fmt.Fprintln(w, "MemoryLimit:", appEnv.MemoryLimit)
		fmt.Fprintln(w, "WorkingDir:", appEnv.WorkingDir)
		fmt.Fprintln(w, "TempDir:", appEnv.TempDir)
		fmt.Fprintln(w, "User:", appEnv.User)

		// Here we check for mysql
		mysql, err := appEnv.Services.WithName("mysql")
		if err != nil {
			fmt.Fprintln(w, "\nMYSQL: false")
		} else {
			fmt.Fprintln(w, "\nMYSQL: true")
			fmt.Fprintln(w, "\nDB Host: ", mysql.Credentials["hostname"])
			fmt.Fprintln(w, "DB Port: ", mysql.Credentials["port"])
			fmt.Fprintln(w, "DB Name: ", mysql.Credentials["name"])
			fmt.Fprintln(w, "DB User: ", mysql.Credentials["username"])
			fmt.Fprintln(w, "DB Pass: ", mysql.Credentials["password"])
		}

		// Here we check for rabbitmq
		rabbitmq, err := appEnv.Services.WithName("rabbitmq")
		if err != nil {
			fmt.Fprintln(w, "\nRABBITMQ: false")
		} else {

			protocolsi, _ := rabbitmq.Credentials["protocols"]
			protocols, _ := protocolsi.(map[string]interface{})
			amqpi, _ := protocols["amqp"]
			amqp, _ := amqpi.(map[string]interface{})
			mgmti, _ := protocols["management"]
			mgmt, _ := mgmti.(map[string]interface{})

			fmt.Fprintln(w, "\nRABBITMQ: true")
			fmt.Fprintln(w, "\nAMQP: ")
			fmt.Fprintln(w, "\thost: ", amqp["host"])
			fmt.Fprintln(w, "\tport: ", amqp["port"])
			fmt.Fprintln(w, "\tuser: ", amqp["username"])
			fmt.Fprintln(w, "\tpass: ", amqp["password"])
			fmt.Fprintln(w, "\tvhost: ", amqp["vhost"])
			fmt.Fprintln(w, "\tssl: ", amqp["ssl"])
			fmt.Fprintln(w, "\nMGMT: ")
			fmt.Fprintln(w, "\thost: ", mgmt["host"])
			fmt.Fprintln(w, "\tport: ", mgmt["port"])
			fmt.Fprintln(w, "\tuser: ", mgmt["username"])
			fmt.Fprintln(w, "\tpass: ", mgmt["password"])
			fmt.Fprintln(w, "\tssl: ", mgmt["ssl"])
		}

		// Here we check for redis
		redis, err := appEnv.Services.WithName("redis")
		if err != nil {
			fmt.Fprintln(w, "\nREDIS: false")
		} else {
			fmt.Fprintln(w, "\nREDIS: true")
			fmt.Fprintln(w, "\nHost: ", redis.Credentials["host"])
			fmt.Fprintln(w, "Pass: ", redis.Credentials["password"])
		}

	} else {
		fmt.Fprintln(w, "Bahh... I'm just running on a normal pc.")
	}

}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v\n", DefaultPort)
		port = DefaultPort
	}

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}
}
