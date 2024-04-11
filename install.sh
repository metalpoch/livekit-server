#!/bin/sh

echo "installing  go"
pkg install -y go

echo "downloading livekit v1.5.3"
cd /tmp
wget https://github.com/livekit/livekit/archive/refs/tags/v1.5.3.tar.gz
tar -xf v1.5.3.tar.gz && cd livekit-1.5.3

echo "compiling go binary"
CGO_ENABLE=0 GOARCH=amd64 GO111MODULE=on go build -a -o ~/livekit-server ./cmd/server

echo "deleting temporal files"
rm /tmp/{livekit-1.5.3,v1.5.3.tar.gz}

echo "~/livekit-server created successfully"

echo "downloading livekit-server-example"
wget https://github.com/metalpoch/livekit-server/archive/refs/tags/deploy.tar.gz
tar -xf deploy.tar.gz && cd livekit-server-deploy/
go mod tidy
CGO_ENABLE=0 GOARCH=amd64 GO111MODULE=on go build -a -o ~/streaming-server main.go

cd
echo "~/streaming-server created successfully"

echo -e 'export LIVEKIT_HOST_URL="http://localhost:7880"\nexport LIVEKIT_API_KEY="devkey"\nexport LIVEKIT_API_SECRET="secret"' >> ~/.profile

export LIVEKIT_HOST_URL="http://localhost:7880"
export LIVEKIT_API_KEY="devkey"
export LIVEKIT_API_SECRET="secret"
