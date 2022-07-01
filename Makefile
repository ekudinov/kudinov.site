.PHONY:build
build:
	rm -Rf build
	mkdir build
	GOARCH=wasm GOOS=js go build -o build/web/app.wasm
	go build -o build/server
	cp -Rf assets build/web

.PHONY:run
run: build
	cd build;./server &

.PHONY:clean
clean:
	rm -Rf build
