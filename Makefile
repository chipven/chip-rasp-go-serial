all:
	cd main && env GOOS=linux GOARCH=arm GOARM=6 go build main.go && scp main pi@192.168.3.98:~