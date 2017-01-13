Go echo server
=======

go-echo is a small HTTP server that **print** and **echo** the request headers and body back as JSON.

## Build it
```
docker build -t flin/go-echo .
```

## Run it
```
docker run -d -p 8000:8000 flin/go-echo
```
