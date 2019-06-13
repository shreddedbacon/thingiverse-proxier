# Thingiverse Proxier

A way to use GO web server to display results from thingiverse api using thingiverse-go

# Run it
```
# edit .env using .env-example as template
# source it
. .env
# get imports
go get -v
# run but don't build
go run *.go
```

# Build it
```
# get imports
go get -v
# build it
go build -o thingiverse *.go
#
chmod +x thingiverse
# source env vars
. .env
# run it
./thingiverse
```

# Use
Once running, open browser for one of the following and click a thing

* http://localhost:8900/newest
* http://localhost:8900/popular
* http://localhost:8900/featured
