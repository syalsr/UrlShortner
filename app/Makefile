run:
	GOOS=linux go build -a -tags netgo -o app cmd/*.go
	docker build -t syalsr .
	docker run -p 8080:8080 syalsr