#!/bin/bash

# cleanup the profile directory before starting
rm -fr profile

# profile each supported package
for what in "list" "progress" "table"
do
    echo "Profiling ${what} ..."
    mkdir -p profile/${what}
    go build -o profile/${what}/${what} cmd/profile-${what}/profile.go
    (cd profile/${what} && \
        ./${what} && \
        go tool pprof -pdf ${what} cpu.pprof > ../${what}.cpu.pdf && \
        go tool pprof -pdf ${what} mem.pprof > ../${what}.mem.pdf)
    echo "Profiling ${what} ... done!"
    echo
done

ls -al profile/*.pdf

