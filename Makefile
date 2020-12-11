build:
	rm -rf target
	mkdir target
	cp cmd/mini/mini-example.yaml target/mini.yaml
	go build -o target/mini cmd/mini/main.go

run:
	nohup target/mini -conf=target/mini.yaml -env=release 2>&1 > target/mini.log &

stop:
	pkill -f target/mini
