BIN=go-zabbix-slack-alert

build:
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	gox -osarch "linux/amd64 linux/386 linux/arm darwin/amd64" -output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}" && \
		mkdir -p distpkg && \
		for ARCH in `ls dist/`; do zip -j -o distpkg/$(BIN)_$${ARCH}.zip dist/$${ARCH}/$(BIN)*; done

clean:
	rm -rf dist/*
	rm -rf distpkg/*

.PHONY: build clean
