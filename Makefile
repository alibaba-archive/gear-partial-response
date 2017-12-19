test:
	go test -v

cover:
	@rm -rf *.coverprofile
	go test -coverprofile=partial.coverprofile -v
	gover
	go tool cover -html=partial.coverprofile
