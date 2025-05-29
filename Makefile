# --- Variables ---
APP_NAME := ward-stock-backend
MAIN := cmd/server/main.go
SEED := cmd/seed/main.go
SWAG_OUT := cmd/server/docs

run:
	@echo "🚀 Running $(APP_NAME)..."
	go run $(MAIN)

seed:
	@echo "🌱 Seeding database..."
	go run $(SEED)

swag:
	swag init --generalInfo "main.go" --output docs

clean-docs:
	rm -rf $(SWAG_OUT)/*