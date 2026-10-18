.PHONY: test progress

# make test            every exercise
# make test EX=1.3.2   one exercise
# make test EX=1.3     one chapter
test:
ifeq ($(EX),)
	go vet ./... && go test ./...
else
	@ex="$(EX)"; case "$$ex" in *.*.*) ch="$${ex%.*}" ;; *) ch="$$ex" ;; esac; tag=$$(echo $$ch | tr . _); \
	case "$$ex" in \
	  *.*.*) go test -count=1 -run "^TestEx$$(echo "$$ex" | tr . _)$$" ./ch$$tag ;; \
	  *) go test -count=1 ./ch$$tag ;; \
	esac
endif

progress:
	python3 tools/progress.py
