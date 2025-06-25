# k8sController

To create a TLS cert

* `openssl req -newkey rsa:2048 -nodes -keyout tls.key -x509 -days 365 -out tls.crt -subj "/CN=image-validator.default.svc"`

* `kubectl create secret tls image-validator-tls --cert=tls.crt --key=tls.key -n default`