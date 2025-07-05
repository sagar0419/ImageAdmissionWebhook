package main

import (
	"fmt"
	"net/http"

	controller "github.com/sagar0419/ImageAdmissionWebhook/api/validate"
)

func main() {
	fmt.Println("Starting Image validation controller")

	http.HandleFunc("/validate", controller.AdmissionRouter)
	err := http.ListenAndServeTLS(":443", "./tls/tls.crt", "./tls/tls.key", nil)
	if err != nil {
		panic(err)
	}
}
