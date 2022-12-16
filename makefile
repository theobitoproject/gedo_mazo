.PHONY: test_all
test_all:
	ginkgo -r --race -cover --randomize-all --randomize-suites --trace -skip-package=vendor

.PHONY: test_all_vv
test_all_vv:
	ginkgo -r --race -cover --randomize-all --randomize-suites --trace -skip-package=vendor --vv

.PHONY: gen_mocks
gen_mocks:
	sh scripts/gen_mocks.sh

.PHONY: gen_gedo_mazo_grpc_stubs
gen_gedo_mazo_grpc_stubs:
	sh scripts/gen_proto.sh gedo_mazo $(version) internal/ports/gedo_mazo
