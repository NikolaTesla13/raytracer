SRC_FILES := $(wildcard src/*.go)

all: run

run: $(SRC_FILES)
	go run $(SRC_FILES)

build: $(SRC_FILES)
	go build -buildmode=exe $(SRC_FILES)

clean: 
	-rm -f *.png
	-del *.png