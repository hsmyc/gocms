run: generate-templates watch-app

generate-templates:
	@dirs=$$(find . -name '*.templ' -exec dirname {} \; | uniq); \
	for dir in $$dirs; do \
		if [ -d "$$dir" ]; then \
			find "$$dir" -name '*.go' -exec rm -f {} \; ; \
			(cd "$$dir" && templ generate); \
		fi \
	done
watch-app:
	@tsc ./static/script/script.ts --outDir ./static/script
	@find . -name '*.templ' | entr -r -s 'make generate-templates & go run cmd/main.go' &
stop:
	@pkill -f 'go run cmd/main.go'
	@pkill -f 'entr'
