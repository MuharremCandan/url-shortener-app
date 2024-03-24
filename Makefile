dev: 
	docker compose up -d && air

hot:
	docker build -t url-shortner:latest .

.PHONY: dev hot