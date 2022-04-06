all: dist rpm

codegen:
	docker build -f hack/Dockerfile . | tee /dev/tty | tail -n1 | cut -d' ' -f3 | xargs -I{} \
		docker run --rm -v $$PWD:/workspace -w /workspace {} hack/codegen.sh

test:
	go test ./... -coverprofile cover.out

dist:
	docker build -f build/dist/Dockerfile . | tee /dev/tty | tail -n1 | cut -d' ' -f3 | xargs -I{} \
		docker run --rm -v $$PWD:/workspace -w /workspace {} build/dist/build.sh

rpm:
	docker build -f build/rpm/Dockerfile . | tee /dev/tty | tail -n1 | cut -d' ' -f3 | xargs -I{} \
		docker run --rm -v $$PWD:/workspace -w /workspace {} build/rpm/build.sh

clean:
	rm -rf out
