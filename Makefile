amd64:
	CGO_ENABLED=false GOOS=linux GOARCH=amd64 go build -o ./zarf/deploy/amd64/app ./cmd
	docker build --platform=amd64 -t digitalcircle/khelloworld:amd64 -f ./zarf/deploy/amd64/Dockerfile .
	docker push digitalcircle/khelloworld:amd64

arm64:
	CGO_ENABLED=false GOOS=linux GOARCH=arm64 go build -o ./zarf/deploy/amd64/app ./cmd
	docker build --platform=arm64 -t digitalcircle/khelloworld:arm64 -f ./zarf/deploy/arm64/Dockerfile .
	docker push digitalcircle/khelloworld:arm64

pub: amd64 arm64