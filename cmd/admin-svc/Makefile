build:
	rm -rf admin-svc && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin-svc && scp admin-svc chries@44.252.33.227:/home/chries-svc && rm -rf admin-svc
