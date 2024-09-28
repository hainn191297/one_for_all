#
MODULE_NAME=one_for_all
SERVICE=services
#
ROOT_DIR=$(shell pwd)
LOGS_DIR=$(ROOT_DIR)/logs
SHARED_DIR=$(ROOT_DIR)/shared
#
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
LOGS_DIR=$(ROOT_DIR)/logs
SCHEMA_DIR=$(ROOT_DIR)/schema
API_DIR=$(SHARED_DIR)/api
# COMMON_DIR=$(ROOT_DIR)/common

# rpc model
PROTO_MODEL_DIR=$(SHARED_DIR)/protobuf

# account api + rpc
ACCOUNT_DIR=$(SERVICE_DIR)/account
ACCOUNT_API_DIR=$(ACCOUNT_DIR)/api
ACCOUNT_MODEL_DIR=$(ACCOUNT_DIR)/model
ACCOUNT_RPC_DIR=$(ACCOUNT_DIR)/rpc
ACCOUNT_RPC_PB_DIR=$(ACCOUNT_RPC_DIR)/pb

dependency-init:
	go mod init $(MODULE_NAME)
	go mod tidy


# https://go-zero.dev/docs/goctl/api/
# gen api
gen-account-service:
	goctl api go -api $(API_DIR)/account/account.api -dir $(ACCOUNT_API_DIR) -style gozero
	@rm -rf $(ACCOUNT_API_DIR)/etc

gen-account-rpc:
	goctl rpc protoc --proto_path=$(PROTO_MODEL_DIR) $(PROTO_MODEL_DIR)/account/account.proto --go_out=$(ACCOUNT_RPC_PB_DIR) --go-grpc_out=$(ACCOUNT_RPC_PB_DIR) --zrpc_out=$(ACCOUNT_RPC_DIR) -style gozero 
	@rm -rf $(ACCOUNT_RPC_DIR)/etc
	@rm -rf $(ACCOUNT_RPC_DIR)/accountclient

# gen config for multiple services
gen-entity-config:
	protoc --proto_path=$(PROTO_MODEL_DIR) --go_out=$(ROOT_DIR)/common --go-grpc_out=$(ROOT_DIR)/common $(PROTO_MODEL_DIR)/entity/entity.proto



# local
run-account:
	go run $(ACCOUNT_API_DIR)/account.go -f etc/account/account-api.yaml

run-account-rpc:
	go run $(ACCOUNT_RPC_DIR)/account.go -f etc/account/account-rpc.yaml