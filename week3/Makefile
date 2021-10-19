export tag=v1.0

clean:
	rm -rf bin
	docker rmi evan/httpserver:${tag}

build:
	echo "building linux httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/httpserver .

image: build
	echo "building httpserver container"
	docker build -t evan2999/httpserver:${tag} .

push: image
	echo "pushing evan/httpserver"
	docker push evan2999/httpserver:${tag}

