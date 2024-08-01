# Go Commands
tidy:
	go mod tidy
run:
	go run ./cmd/webserver/main.go -port=":3030"
test:
	go test -cover ./...
build:
	bash ./scripts/build.sh

# Templ Commands
templ-watch:
	templ generate -watch
templ-gen:
	templ generate

# Tailwind Css Commands
tw-build:
	@./tailwindcss -i ./web/static/css/input.css -o ./web/static/css/default/style.css
tw-watch:
	@./tailwindcss -i ./web/static/css/input.css -o ./web/static/css/default/style.css --watch
tw-minify:
	@./tailwindcss -i ./web/static/css/input.css -o ./web/static/css/default/style.min.css --minify