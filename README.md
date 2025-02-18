# HLS SERVICES

#### This is services for advertising issue, written in Golang
#### It communicates with [json-server](https://www.npmjs.com/package/json-server) via standart [net-http](https://pkg.go.dev/net/http) package or with [cassandra](https://cassandra.apache.org/) via [cassandra-gocql-driver](https://github.com/apache/cassandra-gocql-driver)

### How to build
```bash
$ make
```
or
```bash
$ go build -o bin/ cmd/ahlc/ahlc.go
$ go build -o bin/ cmd/lhlc/lhlc.go
```
### How to run (params are optional, default should works well)
```bash
$ .\bin\ahlc.exe -dbServerType='cassandra' -dbServerAddr='127.0.0.1:9042' -cacheTTL=20
$ .\bin\lhlc.exe -appAddr='127.0.0.1:8081' -ahlcAddr='127.0.0.1:8080'
```
### Run DB (Apache Cassandra)
```bash
$ sudo apt install cassandra
$ sudo service cassandra start
```
### ALTERNATIVE Run DB (json-server)
```bash
$ npx json-server -w ./db/db.json
```

### Run any static web server, for example:
```bash
$ py -m http.server 8000 -d .\assets\
```
### AHLC endpoints
 **GET** `/auction?sourceID={id}&maxDuration={max_duration}`
 - usage exampe curl:
```bash 
$ curl "http://localhost:8080/auction?sourceID=1&maxDuration=32"
```
#### alternative ffmpeg:
```bash 
$ ffplay "http://localhost:8080/auction?sourceID=1&maxDuration=32" -volume 1
```
### LHLC endpoints
 **POST** `/stitching.m3u8?sourceID={source_id}`
 - usage exampe curl:
 ```bash 
$ curl -X POST "http://localhost:8081/stitching.m3u8?sourceID=1" -H "Content-Type: text/plain" --data-binary '#EXTINF:3.999744,
http://localhost:8000/journey0.ts
#EXTINF:4.000600,
http://localhost:8000/journey1.ts
#EXTINF:3.999567,
http://localhost:8000/journey2.ts
#EXTINF:4.000144,
http://localhost:8000/journey3.ts'
```

> **Limitations:**
>
> - The cache can grow infinitely because of the lack of an eviction mechanism.
> - No m3u8 validation for input for second service
