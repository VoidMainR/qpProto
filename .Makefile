all:	install

install:
	go install ./proto ./jsonpb ./ptypes ./protoc-gen-go

clean:
	go clean ./...

nuke:
	go clean -i ./...

regenerate:
	for d in api srv; do \
		for f in $$d/**/proto/*.proto; do \
			protoc --proto_path=${GOPATH}/src --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
    done