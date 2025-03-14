build:
	go build -o hot-news cmd/main.go
deploy:
	scp hot-news ubuntu@ec2-52-78-87-228.ap-northeast-2.compute.amazonaws.com:/home/ubuntu/hot-news 
	