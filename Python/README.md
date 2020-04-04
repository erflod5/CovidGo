```
    docker build --tag pyserver .
    docker run -it -d --name pyApp -p 5001:5001 pyserver
```