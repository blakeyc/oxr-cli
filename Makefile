package:
	rm -rf ./dist/*
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o ./dist/darwin/oxr .
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/linux/oxr .
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o ./dist/windows/oxr.exe .
