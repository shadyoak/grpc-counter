#!/usr/bin/env bash
set -eu

# Generates the embedded website assets served by this program. Must be 
# run any time the files in 'web/static/...' are updated. Requires 
# installation of the go-bindata-assetfs command line tool.

# ------------------
# go-bindata-assetfs
# ------------------
# https://github.com/elazarl/go-bindata-assetfs

#go-bindata-assetfs -pkg 'web' -o 'web/assets.go' 'web/static/...'
go-bindata-assetfs -pkg "web" "web/static/..."
mv -f 'bindata_assetfs.go' 'web/assets.go'
