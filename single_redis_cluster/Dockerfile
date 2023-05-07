# https://redis.io/docs/management/scaling/
FROM redis:latest

ENV CLUSTER_INSTANCE_COUNT=6
EXPOSE 6000-7000/tcp

RUN set -x &&\
    apt update &&\
    apt install -y procps

COPY start_cluster.sh /

ENTRYPOINT ["/start_cluster.sh"]