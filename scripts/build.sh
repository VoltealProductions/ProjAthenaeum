#!/usr/bin/env bash

version=`git describe --tags HEAD`

package_name=athenaeum

platforms=(
"linux/amd64"
#"linux/arm"
#"linux/arm64"
)

make templ-gen
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

	cpyfldr=./releases/release-$GOOS-$GOARCH/$output_name

  mkdir -p $cpyfldr
	cp -r ./public $cpyfldr
	cp -r .env $cpyfldr

  echo "Building releases/release-$GOOS-$GOARCH-$version/$output_name..."
  env GOOS=$GOOS GOARCH=$GOARCH go build \
  -ldflags "-X github.com/VoltealProductions/athenaeum.Version=$version" \
  -o $cpyfldr/$output_name ./cmd/client/main.go
  if [ $? -ne 0 ]; then
      echo 'An error has occurred! Aborting.'
      exit 1
  fi
  echo "Creating zip archives..."
  zname=$package_name-$version-$GOOS-$GOARCH.zip
  if [ $os = "windows" ]; then
      zip -r ./releases/release-$GOOS-$GOARCH/$zname ./releases/release-$GOOS-$GOARCH/$output_name
      chmod a+x ./releases/release-$GOOS-$GOARCH/$zname
      zip -sf ./releases/release-$GOOS-$GOARCH/$zname
      rm -rf ./releases/release-$GOOS-$GOARCH/$output_name
  else
      zip -r ./releases/release-$GOOS-$GOARCH/$zname ./releases/release-$GOOS-$GOARCH/$output_name
      chmod a+x ./releases/release-$GOOS-$GOARCH/$zname
      zip -sf ./releases/release-$GOOS-$GOARCH/$zname
      rm -rf ./releases/release-$GOOS-$GOARCH/$output_name
  fi
done