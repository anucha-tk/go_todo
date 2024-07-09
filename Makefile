init:
	@echo "== 👩‍🌾 init =="
	brew install pre-commit
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

ci.lint:
	@echo "== 🙆 ci.linter =="
	golangci-lint run -v ./... --fix

dev-templ:
	templ generate -watch -proxy=http://localhost:3000

dev-tailwind:
	npx tailwindcss -i ./internal/assets/tailwind.css -o ./internal/assets/dist/styles.css --watch
