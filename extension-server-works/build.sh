#!/bin/bash

current_dir=$(pwd)

echo "Current Dir IS ${current_dir}"
go build -C ./sql-inline-lsp -tags="prodStd" -o "../../extension-host/dist/server/lsp-server.exe"
echo "build done"
