#!/bin/bash

# Build the binary
go build -o tide ./cmd/tide/main.go

# Move to a directory in PATH
if [ -d "/usr/local/bin" ]; then
    sudo mv tide /usr/local/bin/
    echo "tide has been installed to /usr/local/bin/tide"
else
    echo "Please move the 'tide' binary to a directory in your PATH."
fi
