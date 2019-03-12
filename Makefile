default:
	go get -u github.com/dvyukov/go-fuzz/...
	go get -d github.com/dvyukov/go-fuzz-corpus
	cd hcl
	go-fuzz-build
