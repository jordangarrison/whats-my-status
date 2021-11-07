#!/bin/bash

PACKAGE_NAME=$1
VERSION=$2
fast_flag=$3
if [ -z $fast_flag ]; then
    # platforms=("$(echo {linux,windows,darwin}/{amd64,386})")
    platforms=(linux/amd64 linux/386 windows/amd64 darwin/amd64)
elif [[ "$fast_flag" = "fast" ]]; then
    platforms=("linux/amd64")
else
    echo "Error, invlaid platform specification, exiting"
    exit 1
fi
echo "Building for the following platforms and architectures"
echo ${platforms[@]}

# Vet the code
echo "Vetting code"
go vet
RETVAL=$?
if [ $RETVAL -ne 0 ]; then
    echo "Did not pass \`go vet\`, exitting"
    exit 1
fi

build_platform() {
    local __platform=$1
    echo "Building ${platform}"
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=$PACKAGE_NAME'-'$GOOS'-'$GOARCH'-'$VERSION

    if [[ $GOOS = "windows" ]]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name
    if [ $? -ne 0 ]; then
        echo "An error has occured! Aborting build"
        exit 2
    fi

}

for platform in ${platforms[@]}; do
    build_platform $platform
done
