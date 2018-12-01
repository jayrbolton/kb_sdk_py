#!/bin/bash

platforms=("darwin/amd64" "linux/amd64" "darwin/386" "linux/386" "linux/arm")

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}
  output_name='cli_'$GOOS'_'$GOARCH
  echo "building to ./dist/$output_name"
  env GOOS=$GOOS GOARCH=$GOARCH go build -o dist/$output_name
  if [ $? -ne 0 ]; then
    echo 'An error has occured, aborting.'
    exit 1
  fi
done

echo '..done'
