# Makefile for deploying Kubernetes resources

# Default namespace
NAMESPACE ?= default

# YAML files
DEPLOYMENT_FILE = deployment.yaml
SERVICEACCOUNT_FILE = serviceAccount.yaml
DEPLOYDISALLOW_FILE = disallowedDeploy.yaml
DEPLOYALLOW_FILE = allowedDeploy.yaml
PODDISALLOW_FILE = disallowedPod.yaml
PODALLOW_FILE = allowedpod.yaml
# ROLE_BINDING_FILE = RoleRoleBinding.yaml

# Apply all manifests
.PHONY: deploy
deploy: apply-serviceaccount apply-deployment #apply-rolebinding
	@echo "✅ All resources deployed successfully."

# Apply ServiceAccount (and RBAC if included)
.PHONY: apply-serviceaccount
apply-serviceaccount:
	@echo "🔄 Applying ServiceAccount..."
	kubectl apply -n $(NAMESPACE) -f manifest/$(SERVICEACCOUNT_FILE)

# # Apply Deployment
# .PHONY: apply-rolebinding
# apply-rolebinding:
# 	@echo "🔄 Applying Deployment..."
# 	kubectl apply -n $(NAMESPACE) -f manifest/$(DEPLOYMENT_FILE)

#Apply Deployment
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



# ####################################
#         Testing Webhook            #
######################################

# Apply all test manifests
.PHONY: test-deploy
test-deploy: test-deployment test-pod
	@echo "✅ Deployment and Test Deployed"


# Test Deployment
.PHONY: test-deployment
test-deployment:
	@echo "🔄 apply-test-deployment..."
	kubectl apply -n $(NAMESPACE) -f manifest/testManifests/deployment/$(DEPLOYDISALLOW_FILE)

# Test Pod
.PHONY: test-pod
test-pod:
	@echo "🔄 apply-test-deployment..."
	kubectl apply -n $(NAMESPACE) -f manifest/testManifests/pod/$(PODDISALLOW_FILE)

# Delete Test
.PHONY: test-clean
test-clean: 
	@echo "☠️ Deleting all test..."
	-kubectl delete -n $(NAMESPACE) -f manifest/testManifests/deployment/$(DEPLOYDISALLOW_FILE)
	-kubectl delete -n $(NAMESPACE) -f manifest/testManifests/pod/$(PODDISALLOW_FILE)

# check test status
.PHONY: test-status
test-status:
	@echo "☠️ Getting test status..."
	kubectl get all -n $(NAMESPACE) | grep disallowed 
