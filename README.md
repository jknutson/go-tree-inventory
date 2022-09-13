# go-tree-inventory

A simple Go app and HTML page to faciliate the recording of tree inventory data.

## Quickstart

### SSL

SSL is required for geolocation to work. You can create a self-signed cert with the command below:

```sh
mkdir -p ssl && cd ssl
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes
```

### SQL

See the `sql/` folder for scripts to create the table to hold the `POST`ed data.

There is code for an optional view as well that can be used by GIS applications.

### Go HTTP Server

To start the Go HTTP server:

```sh
go run main.go
```
