#mises_alpha
truss:
	truss proto/vpnsvc.proto  --pbpkg github.com/mises-id/mises-vpnsvc/proto --svcpkg github.com/mises-id/mises-vpnsvc --svcout . -v
run:
	go run cmd/cli/main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/cli/main.go
upload:
	scp ./main mises_alpha:/apps/mises-vpnsvc/
replace:
	ssh mises_alpha "mv /apps/mises-vpnsvc/main /apps/mises-vpnsvc/mises-vpnsvc"
restart:
	ssh mises_alpha "sudo supervisorctl restart mises-vpnsvc"
deploy: build \
	upload \
	replace \
	restart
#mises_backup
upload-backup:
	scp ./main mises_backup:/apps/mises-vpnsvc/
replace-backup:
	ssh mises_backup "mv /apps/mises-vpnsvc/main /apps/mises-vpnsvc/mises-vpnsvc"
restart-backup:
	ssh mises_backup "sudo supervisorctl restart mises-vpnsvc"
deploy-backup: build \
	upload-backup \
	replace-backup \
	restart-backup 
#mises_master
upload-master:
	scp ./main mises_master:/apps/mises-vpnsvc/
replace-master:
	ssh mises_master "mv /apps/mises-vpnsvc/main /apps/mises-vpnsvc/mises-vpnsvc"
restart-master:
	ssh mises_master "sudo supervisorctl restart mises-vpnsvc"
deploy-master: build \
	upload-master \
	replace-master \
	restart-master 
	