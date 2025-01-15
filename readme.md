## Suiren, the Pygmy Water Lily

### Stack

`net/http` `chi` `pgx` `golang-migrate`

### Dev Notes

```bash
# Install Dependencies
go mod tidy

# Install golang-migrate

# Linux / Direct Install
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# MacOS (brew)
brew install golang-migrate

# Windows (screw you)
scoop install migrate

# Create Migration
make migrate-up

# Run the server
make run
```