## Suiren, the Pygmy Water Lily

### Stack

`chi` `pgx` `golang-migrate`

### Dev Notes

MacOS
```bash
# Install Dependencies
go mod tidy

# Install golang-migrate
brew install golang-migrate

# Create Migration
make migrate-up

# Run the server
make run
```