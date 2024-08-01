# Go Commands
tidy:
	go mod tidy
run:
	go run ./cmd/athenaeum/main.go -domain=":3030"
test:
	go test -cover ./athenaeum/...
build:
	go build -v -o ./bin/athenaeum/app ./cmd/athenaeum/main.go
br:
	bash ./scripts/build-release.sh

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