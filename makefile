.PHONY: test_all
test_all:
	ginkgo -r --race -cover --randomize-all --randomize-suites --trace -skip-package=vendor

.PHONY: test_all_vv
test_all_vv:
	ginkgo -r --race -cover --randomize-all --randomize-suites --trace --vv -skip-package=vendor
