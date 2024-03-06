run: generate-templates watch-app browser-sync

generate-templates:
	@dirs=$$(find . -name '*.templ' -exec dirname {} \; | uniq); \
	for dir in $$dirs; do \
		if [ -d "$$dir" ]; then \
			find "$$dir" -name '*.go' -exec rm -f {} \; ; \
			(cd "$$dir" && templ generate); \
		fi \
	done

watch-app:
	@find . -name '*.templ' | entr -r -s 'make generate-templates & go run cmd/main.go' &
	@find . -name '*.go' | entr -r -s 'go run cmd/main.go' &
	@find . -name '*.ts' | entr -r -s 'go run cmd/main.go' &
	@find . -path ./static/dist -prune -o -name '*.css' -print | entr -r -s 'go run cmd/main.go' &

browser-sync:
	@browser-sync start --proxy "localhost:8080" --files "static/dist/**/, views/**/, *.go"

stop:
	@pkill -f 'go run cmd/main.go'
	@pkill -f 'entr'
	@pkill -f 'browser-sync'
