build:
	rm -rf rszai-svc && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rszai-svc

push:
	scp rszai-svc chries@44.252.33.227:/home/chries/rszai-svc && rm -rf rszai-svc

