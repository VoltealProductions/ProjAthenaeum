# Podman Commands
docker-up:
	@docker compose --env-file .env build && docker compose --env-file .env up
docker-down:
	@docker compose --env-file .env down

# Templ Commands
templ-watch:
	@templ generate -watch
templ-gen:
	@templ generate

# Tailwind Css Commands
tw-build:
	@./tailwindcss -i ./public/css/input.css -o ./public/css/default/style.css
tw-watch:
	@./tailwindcss -i ./public/css/input.css -o ./public/css/default/style.css --watch
tw-minify:
	@./tailwindcss -i ./public/css/input.css -o ./public/css/default/style.min.css --minify

# Go Commands
tidy:
	@go mod tidy
run: templ-gen tw-build
	@go run ./cmd/client/main.go
test:
	@go test -cover ./...

# Bash Scripts
build-release:
	@bash ./scripts/build.sh