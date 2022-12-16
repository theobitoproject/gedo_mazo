#!/bin/bash

readonly SERVICE=$1
readonly VERSION=$2
readonly TARGET_PATH=$3
readonly TMP_PATH="tmp"

# common functions

check_empty() {
  if [ -z "$2" ]
  then
    print_error "$1 cannot be empty"
    exit
  fi
}

get_tar_gz_filename() {
  local filename=$1
  local pre="${filename%.*}"
  echo "${pre%.*}"
}

print_error() {
  echo "\nERROR: $1"
}

print_title() {
  local title=$1
  echo "\n${title}"
}

# common functions

print_title "gen proto started..."

# checking arguments
check_empty SERVICE $SERVICE
check_empty VERSION $VERSION
check_empty TARGET_PATH $TARGET_PATH
# checking arguments

print_title "pulling protos..."

rm $TMP_PATH/*.tar.gz

LINK="https://github.com/theobitoproject/toad_contracts/archive/refs/tags/${VERSION}.tar.gz"

curl -LJO --output-dir $TMP_PATH --fail $LINK

CURL_ERROR=$?
if [ $CURL_ERROR -ne 0 ]
then
  print_error "failed downloading files from repo. curl error code: ${CURL_ERROR}"
  exit
fi

print_title "protos downloaded"

print_title "extracting protos..."

FILENAME="$(ls $TMP_PATH | egrep '\.tar.gz$' | head -n 1)"
EXTRACT_PATH="$TMP_PATH/$(get_tar_gz_filename $FILENAME)"

tar -xf $TMP_PATH/$FILENAME -C $TMP_PATH

print_title "protos extracted"

print_title "proto stub generation started..."

mkdir $TARGET_PATH

PROTO_PATH="$EXTRACT_PATH/proto/$SERVICE"

protoc \
  "-I=$PROTO_PATH" $PROTO_PATH/*.proto \
  "--go_out=$TARGET_PATH" --go_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "--go-grpc_out=$TARGET_PATH" --go-grpc_opt=paths=source_relative

print_title "proto stub generation finished"

print_title "cleaning..."

rm $TMP_PATH/*.tar.gz
rm -r $EXTRACT_PATH

print_title "clean was complete"

print_title "gen proto finished"
