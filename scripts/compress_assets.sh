#!/usr/bin/env bash

# https://misterorion.com/caddy-server-brotli/
find ./dist -type f -size +1400c \
    -regex ".*\.\(css\|html\|js\|json\|svg\|xml\)$" \
    -exec brotli --best -f {} \+ \
    -exec gzip --best -k -f {} \+ \
    -exec zstd -19 -f {} \+

