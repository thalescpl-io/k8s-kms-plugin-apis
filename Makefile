.PHONY: all gen


## Dev
gen: gen-grpc
gen-grpc:
		@mkdir -p generated
		@prototool all || true
		@cp -r generated/github.com/thalescpl-io/k8s-kms-plugin-apis/* .
		@cp -r generated/istio/* istio/
		@cp -r generated/k8s/* k8s/
		@cp -r generated/simplekms/* simplekms/
		@rm -rf generated/

