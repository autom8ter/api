.PHONY: proto
proto: ## generate protobufs
	@docker run -v `pwd`:/defs namely/prototool:1.17_0 generate

clean:
	@rm -rf sdk/*

.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'