### Docker

```sh
docker build -t stresscli .
```

### Running

```sh
docker run stresscli stress --url=http://google.com --requests=1000 --concurrency=10
```
