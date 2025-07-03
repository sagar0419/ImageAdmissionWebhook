# k8sController

## To create a TLS certs 🔒 use follwing command

```

openssl genrsa -out ca.key 2048

openssl req -new -x509 -days 365 -key ca.key \
  -subj "/C=AU/CN=image-validation-controller.default.svc"\
  -addext "subjectAltName = DNS:image-validation-controller.default.svc" \
  -out ca.crt

openssl req -newkey rsa:2048 -nodes -keyout server.key \
  -subj "/C=AU/CN=image-validation-controller.default.svc" \
  -addext "subjectAltName = DNS:image-validation-controller.default.svc" \
  -out server.csr

openssl x509 -req \
  -extfile <(printf "subjectAltName=DNS:image-validation-controller.default.svc,DNS:image-validation-controller.default.svc.cluster.local") \
  -days 365 \
  -in server.csr \
  -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out server.crt

  kubectl create secret tls simple-kubernetes-webhook-tls \
  --cert=ca.crt \
  --key=ca.key \
  --dry-run=client -o yaml \
  > ./secret.yaml


cat ca.crt | base64 | tr -d '\n'

rm ca.crt ca.key ca.srl server.crt server.csr server.key
```

## Usage ⚙️

Run these commands in the [Makefile](./Makefile) directory 🗂️:

`make deploy          # Apply ServiceAccount + Deployment`

`make status          # See what’s running`

`make clean           # Delete all deployed resources`

You can override the namespace if needed:

`make deploy NAMESPACE=kube-system`


If you are cloning the code in the local machine and running it from there use [localValidation.yaml](manifest/localValidation.yaml) file.

If you are running it on kubernetes than use this validation file [validateWebhookCOnfiguration.yaml](manifest/validateWebhookCOnfiguration.yaml) .

To test this validation controller I have added the manifest file here [testManifests](manifest/testManifests).