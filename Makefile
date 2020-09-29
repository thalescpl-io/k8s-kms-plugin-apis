.PHONY: all lint build coverage dev  gen


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
#gen-openapi:
#		@swagger generate server --quiet -m pkg/est/models -s pkg/est/restapi -f apis/kms/v1/est.yaml
#		@swagger generate client --quiet --existing-models=pkg/est/models -c pkg/est/client -f apis/kms/v1/est.yaml
