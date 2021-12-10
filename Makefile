build:
	go build -o src main.go

run: build 
	./server

watch:
	reflex -s -r '\.go$$' make run