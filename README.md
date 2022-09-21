# redis-checkhealth

## Usage

Compile

```bash
go get -d .
go build redis-checkhealth.go
```

Run the redis-checkhealth

```bash
./redis-checkhealth:7550 <REDIS_IP>:<REDIS_PORT> <REDIS_PASSWORD>
```

NOAUTH

```bash
./redis-healthcheck:7550 127.0.0.1:6379 ""
```

With AUTH

```bash
./redis-checkhealth:7550 127.0.0.1:6379 xyz123
OR
./redis-checkhealth:7550 127.0.0.1:6379 "xyz123"
```


Install On Linux

```bash
edit the redis-ch.conf, configure variable redis server, port and redis password
./install.sh

```





## Test

```bash
curl http://localhost:1323/healthz
```



## Response

200 (OK)

```bash
Redis OK
```

500 (Error)

```bash
Redis ERROR
```


