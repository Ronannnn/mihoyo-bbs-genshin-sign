# Mihoyo BBS Genshin Impact Sign Tool

This is a tool to sign for Mihoyo Chinese BBS Genshin Impact.

## How to use

```bash
docker run -d -p 5001:5001 ronannnn/mihoyo-bbs-genshin-sign:latest
```

Or you can download [docker-compose.yaml](https://github.com/Ronannnn/mihoyo-bbs-genshin-sign/blob/main/docker-compose.yaml)
and
```bash
docker-compose up -d
```

## BBS Service

See more details in `service/sign.go`

## API

[APIs](https://editor.swagger.io/?url=https://raw.githubusercontent.com/Ronannnn/mihoyo-bbs-genshin-sign/main/swagger.yaml)
