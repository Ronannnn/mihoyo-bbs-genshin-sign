# Mihoyo BBS Genshin Impact Sign Tool

This is a tool to sign for Mihoyo Chinese BBS Genshin Impact.

## How to use

```bash
docker run -d \
-v $(pwd)/sign_config:/app/sign_config \
-v $(pwd)/sign_data:/app/sign_data \
-v $(pwd)/sign_log:/app/sign_log \
-p 5001:5001 \
-p 9900:9900 \
ronannnn/mihoyo-bbs-genshin-sign:latest
```

Or you can download [docker-compose.yaml](https://github.com/Ronannnn/mihoyo-bbs-genshin-sign/blob/main/deployments/docker-compose/docker-compose.yaml)
and
```bash
docker-compose up -d
```

## BBS Service

See more details in `internal/service/sign.go`

## API

[APIs](https://editor.swagger.io/?url=https://raw.githubusercontent.com/Ronannnn/mihoyo-bbs-genshin-sign/main/api/swagger.yaml)
