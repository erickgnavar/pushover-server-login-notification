build:
	go build -ldflags "-X main.pushover_app_token=$(PUSHOVER_APP_TOKEN) -X main.pushover_user_key=$(PUSHOVER_USER_KEY)" main.go

build_linux:
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.pushover_app_token=$(PUSHOVER_APP_TOKEN) -X main.pushover_user_key=$(PUSHOVER_USER_KEY)" main.go

build_arm:
	GOARCH=arm GOOS=linux GOARM=5 go build -ldflags "-X main.pushover_app_token=$(PUSHOVER_APP_TOKEN) -X main.pushover_user_key=$(PUSHOVER_USER_KEY)" main.go
