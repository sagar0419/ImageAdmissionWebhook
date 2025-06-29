# Makefile for deploying Kubernetes resources

# Default namespace
NAMESPACE ?= default

# YAML files
DEPLOYMENT_FILE = deployment.yaml
SERVICEACCOUNT_FILE = serviceAccount.yaml
ROLE_BINDING_FILE = RoleRoleBinding.yaml

# Apply all manifests
.PHONY: deploy
deploy: apply-serviceaccount apply-deployment apply-rolebinding
	@echo "✅ All resources deployed successfully."

# Apply ServiceAccount (and RBAC if included)
.PHONY: apply-serviceaccount
apply-serviceaccount:
	@echo "🔄 Applying ServiceAccount..."
	kubectl apply -n $(NAMESPACE) -f manifest/$(SERVICEACCOUNT_FILE)

# Apply Deployment
.PHONY: apply-rolebinding
apply-rolebinding:
	@echo "🔄 Applying Deployment..."
	kubectl apply -n $(NAMESPACE) -f manifest/$(DEPLOYMENT_FILE)

# Apply Deployment
.PHONY: apply-deployment
apply-deployment:
	@echo "🔄 Applying Role and RoleBinding..."
	kubectl apply -n $(NAMESPACE) -f manifest/$(ROLE_BINDING_FILE)

# Delete all resources
.PHONY: clean
clean:
	@echo "☠️ Deleting resources..."
	-kubectl delete -n $(NAMESPACE) -f manifest/$(DEPLOYMENT_FILE)
	-kubectl delete -n $(NAMESPACE) -f manifest/$(SERVICEACCOUNT_FILE)
	-kubectl delete -n $(NAMESPACE) -f manifest/$(ROLE_BINDING_FILE)

# Check status
.PHONY: status
status:
	@echo "☠️ Getting resource status..."
	kubectl get all -n $(NAMESPACE)
