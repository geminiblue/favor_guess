COMPILE_TIME = $(shell date +"%Y%M%d%H%M%S")
GIT_REVISION = $(shell git show -s --pretty=format:%h)
CFLAGS += "\"$(GIT_REVISION)\""
build:
	go clean
	rm -rf bin/*
	go build -o bin/fbwg_drawin_$(COMPILE_TIME)_${GIT_REVISION}
build_linux:
	GOARCH=amd64 GOOS=linux go build -o bin/fbwg_linux_amd64_${COMPILE_TIME}_$(GIT_REVISION)
clean:
	go clean
	rm -rf bin/*
