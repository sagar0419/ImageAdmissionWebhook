# k8sController

## To create a TLS cert 🔒

* `openssl req -newkey rsa:2048 -nodes -keyout tls.key -x509 -days 365 -out tls.crt -subj "/CN=image-validator.default.svc"`

* `kubectl create secret tls image-validator-tls --cert=tls.crt --key=tls.key -n default`

## Usage ⚙️

Run these commands in the Makefile directory 🗂️:

`make deploy          # Apply ServiceAccount + Deployment
make status          # See what’s running
make clean           # Delete all deployed resources`

You can override the namespace if needed:

`make deploy NAMESPACE=kube-system`