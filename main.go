package main

import (
	"fmt"
	"net/http"

	controller "github.com/sagar0419/k8sController/controller"
)

func main() {
	fmt.Println("Starting Image validation controller")

	http.HandleFunc("/validate", controller.AdmitPods)
	err := http.ListenAndServeTLS(":8443", "./tls/tls.crt", "./tls/tls.key", nil)
	// err := http.ListenAndServeTLS(":8443", "/tls/tls.crt", "/tls/tls.key", nil)
	if err != nil {
		panic(err)
	}
}
