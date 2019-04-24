.PHONY: proto
proto: ## generate protobufs
	@docker run -v `pwd`:/defs namely/prototool:1.17_0 generate

.PHONY: def
def:  ## generate protobuf descriptor file
	@docker run -v `pwd`:/defs namely/protoc:1.17_0 \
		-I . \
        --include_imports \
        --include_source_info \
        --descriptor_set_out descriptor.pb \
        api.proto

.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'