DB_URL="postgres://postgres:1@localhost:5432/test?sslmode=disable"
mi_up:
	migrate -database $(DB_URL) -path "migrations/" up
mi_down:
	migrate -database $(DB_URL) -path "migrations/" down
mi_force:
	migrate -database $(DB_URL) -path "migrations/" force $(v)