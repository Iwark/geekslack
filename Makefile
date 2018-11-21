build:
	GOOS=linux go build -o bin/geekslack cmd/geekslack/*.go

deploy: build
	zip -j bin/geekslack.zip bin/geekslack

.PHONY: build deploy
