#!/bin/sh

minorVer=$(cat "_run/minor-ver")

minorVer=$((minorVer + 1))
echo "$minorVer" > "_run/minor-ver"

echo "1" > "_run/build-ver"

exit 0

