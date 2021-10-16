# Action
- `pause`
- `unpause`
- `empty`
- `info`
- `check`

# Worker Pool
- `1 <= n <= len(target)` 
- `0` for unlimited pool depend 

# Target
    Array of topics or channels
### Example
- topic: `topicsatu`
- channel: `topicsatu/channelsatu`

# Development
```$ docker-compose up -d``` 

# Run
- ```$ go build```
- ```$ ./nsqpause``` 
```
╰─$ ./nsqpause
nsqpause is ready.
topicsatu/channeldua is paused
topicsatu/channelsatu is paused
topicdua is paused
```