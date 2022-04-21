BUILD_NAME=es_bulk
SOURCE=main.go

.PHONY: build clean 
default: build

build:
	go build -o ${BUILD_NAME}.app/Contents/MacOS/${BUILD_NAME} ${SOURCE}

clean:
	rm -rf ./${BUILD_NAME}.app
