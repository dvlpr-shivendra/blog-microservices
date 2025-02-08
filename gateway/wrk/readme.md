# wrk Load Testing

### Install wrk
```brew install wrk  # macOS```

```sudo apt install wrk  # Ubuntu/Debian```

### Stress test an endpoint
```wrk -t4 -c1000 -d30s -s post.lua http://localhost:8080/api/posts/2/likes```

**post.lua**
```
wrk.method = "POST"
wrk.body = '{"title": "Hello world"}'
wrk.headers["Content-Type"] = "application/json"
```