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