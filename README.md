# Twitter tracker 🦆
> Blazing-fast tweets scraper without any API Keys!

## Installation
```bash
docker run --name twitter-tracker -p 3000:3000 -d ghcr.io/matisiekpl/twitter-tracker:master
```

## Usage
For query: `q=elon musk` and `n=100`
```
GET /search?q=elon%20musk&n=100
```

## Monitoring
```
GET /metrics
```
