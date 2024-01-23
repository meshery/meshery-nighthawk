check:
	golangci-lint run

protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/meshery/meshery/master/meshes/meshops.proto

proto:	
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

docker:
	DOCKER_BUILDKIT=1 docker build -t layer5/meshery-nighthawk .

docker-run:
	(docker rm -f meshery-nighthawk) || true
	docker run --name meshery-nighthawk -d \
	-p 10000:10000 \
	-e DEBUG=true \
	layer5/meshery-nighthawk:edge-latest

run:
	DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go

run-force-dynamic-reg:
	FORCE_DYNAMIC_REG=true DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go
error:
	go run github.com/meshery/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers

test:
	export CURRENTCONTEXT="$(kubectl config current-context)" 
	echo "current-context:" ${CURRENTCONTEXT} 
	export KUBECONFIG="${HOME}/.kube/config"
	echo "environment-kubeconfig:" ${KUBECONFIG}
	GOPROXY=direct GOSUMDB=off GO111MODULE=on go test -v ./...