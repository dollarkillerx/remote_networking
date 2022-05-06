build_traversal:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o traversal_backend traversal/backend/main.go
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o traversal_agent traversal/agent/main.go

build_traversal_upx:
	upx traversal_backend
	upx traversal_agent

build_bulldozer:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o bulldozer_agent bulldozer/cmd/agent/main.go
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o bulldozer_backend bulldozer/cmd/backend/main.go

build_bulldozer_upx:
	upx bulldozer_agent
	upx bulldozer_backend
