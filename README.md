# k8sController

To create a TLS cert

* `openssl genpkey -algorithm RSA -out server.pem -pkeyopt rsa_keygen_bits:2048`

* `openssl req -new -key server.pem  -out server.csr -subj "/CN=validating-webhook-svc.webhooktest.svc" -addext "subjectAltName = DNS:validating-webhook-svc.webhooktest.svc"`

* `openssl req -x509 -sha256 -newkey rsa:2048 -keyout rootCA.pem -out rootCA.crt  -days 650 -subj "/CN=RootCA" -nodes`

* `openssl x509 -copy_extensions copy -req -CA rootCA.crt -CAkey rootCA.pem -in server.csr -out server.crt -days 650 -CAcreateserial`