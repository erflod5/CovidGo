```
docker build -t goserver .
docker run -d --name goApp -p 8001:8001 goserver
```