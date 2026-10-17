.PHONY: test progress

# make test            every exercise
# make test EX=3.2     one exercise
# make test EX=3       one chapter
test:
ifeq ($(EX),)
	go vet ./... && go test ./...
else
	@ch=$$(printf '%02d' $$(echo "$(EX)" | cut -d. -f1)); \
	case "$(EX)" in \
	  *.*) go test -count=1 -run "^TestEx$$(echo "$(EX)" | tr . _)$$" ./ch$$ch ;; \
	  *) go test -count=1 ./ch$$ch ;; \
	esac
endif

progress:
	python3 tools/progress.py
