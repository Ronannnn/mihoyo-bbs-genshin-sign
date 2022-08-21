#!/bin/zsh

image_tag=$1
if [ "$image_tag" = "" ]; then
  echo "Usage: $0 <image_tag>"
  exit 0
fi

image_name_prefix="ronannnn/mihoyo-bbs-genshin-sign"
docker build -t "${image_name_prefix}:${image_tag}" -f Dockerfile .
docker push "${image_name_prefix}:${image_tag}"

# tag as latest and push
docker tag "${image_name_prefix}:${image_tag}" "${image_name_prefix}:latest"
docker push "${image_name_prefix}:latest"
