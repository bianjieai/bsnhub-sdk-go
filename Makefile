PACKAGES=$(shell go list ./...)
export GO111MODULE = on

proto-gen:
	@./third_party/protocgen.sh