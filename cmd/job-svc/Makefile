build:
	rm -rf job-svc && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o job-svc && scp job-svc chries@44.252.33.227:/home/chries/job-svc && rm -rf job-svc
