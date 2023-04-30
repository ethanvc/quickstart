FROM ubuntu/mysql:latest
LABEL authors="testauthor"

RUN set -x && buildDeps='golang-go' \
    && apt update \
    && apt install -y $buildDeps

ENTRYPOINT []