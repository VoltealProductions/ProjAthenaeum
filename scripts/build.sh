#!/usr/bin/env bash

version=`git describe --tags HEAD`

package_name=athenaeum

platforms=(
"linux/amd64"
"linux/arm"
"linux/arm64"
)

make templ-gen
make tw-build
make tw-minify

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  os=$GOOS
  if [ $os = "darwin" ]; then
      os="macOS"
  fi
  output_name="${package_name}-${version}"
  if [ $os = "windows" ]; then
      output_name+='.exe'
  fi

  echo "Building releases/release-${os}-${GOARCH}/$version/$output_name..."
  env GOOS=$GOOS GOARCH=$GOARCH go build \
  -ldflags "-X github.com/VoltealProductions/athenaeum.Version=$version" \
  -o ./releases/release-$GOOS-${GOARCH}/$version/$output_name ./cmd/webserver/main.go
  if [ $? -ne 0 ]; then
      echo 'An error has occurred! Aborting.'
      exit 1
  fi
done