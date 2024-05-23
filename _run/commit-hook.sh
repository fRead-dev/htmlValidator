#!/bin/sh

fileConst="$PWD/const.go"
dateNow=$(date +"%m-%d-%Y")
appName=$(cat "_run/build-name")

##################################################################################
majorVer=$(cat "_run/major-ver")
minorVer=$(cat "_run/minor-ver")
buildVer=$(cat "_run/build-ver")

VERSION="$majorVer.$minorVer.$buildVer"

NAME=$(git branch | grep '*')
echo "$NAME [$VERSION]" ': ' $(cat "$1") > "$1"

###################################

echo "package $appName" > "$fileConst"
echo "" >> "$fileConst"

echo "" >> "$fileConst"
echo "const GlobalVersion string = \"$VERSION\"" >> "$fileConst"
echo "const GlobalDateUpdate string = \"$dateNow\"" >> "$fileConst"
echo "const GlobalName string = \"$appName\"" >> "$fileConst"


#Запись инкремента версии сборки
buildVer=$((buildVer + 1))
echo "$buildVer" > "_run/build-ver"

exit 0

