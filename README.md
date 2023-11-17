# SETUP
Copy `config.sample.env` to `config.env`. Change DBPASS field.

# RUN
`go run .`

# Build
`go build .`

# Docker Run
`docker-compose up --build`


# Local setup and run with air
```
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

alias air='$(go env GOPATH)/bin/air'

air -v

air run .
```


# Copy Only Data Locally 
```
sudo psql -U postgres -d gmv_inventory_nightly -f gmv_inventory_service.sql -h localhost
```

# Copy Full Database Locally 
```
sudo psql -U postgres -d gmv_inventory_nightly -f gmv_inventory_service_full.sql -h localhost
```