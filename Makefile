# --- Variables ---
APP_NAME := ward-stock-backend
MAIN := cmd/server/main.go
SEED := cmd/seed/main.go

run:
	@echo "🚀 Running $(APP_NAME)..."
	go run $(MAIN)

seed:
	@echo "🌱 Seeding database..."
	go run $(SEED)