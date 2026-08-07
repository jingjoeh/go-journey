.PHONY: format check-starters verify-solutions verify-race

LABS := $(sort $(wildcard labs/[0-9][0-9]-*))

format:
	gofmt -w $$(find labs -name '*.go' -type f)

check-starters:
	@for lab in $(LABS); do \
		echo "checking $$lab/starter"; \
		(cd $$lab && go test ./starter) || exit 1; \
	done

verify-solutions:
	@for lab in $(LABS); do \
		echo "testing $$lab/solution"; \
		(cd $$lab && go test ./solution) || exit 1; \
	done

verify-race:
	@for lab in labs/04-concurrency labs/10-caching labs/12-observability labs/14-production-readiness; do \
		echo "race testing $$lab/solution"; \
		(cd $$lab && go test -race ./solution) || exit 1; \
	done
