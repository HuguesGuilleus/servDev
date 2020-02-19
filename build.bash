#!/bin/bash

# HTTP Server for developping project
# 2019 GUILLEUS Hugues <ghugues@netc.fr>
# BSD 3-Clause "New" or "Revised" License

clear
cd templ/ && ./build.bash && cd ..
go build
./servHttp
