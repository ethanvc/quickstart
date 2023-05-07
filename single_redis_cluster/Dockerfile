# https://redis.io/docs/management/scaling/
# docker build -t sjw2docker/single_redis_cluster:7.0.11 .
FROM redis:7.0.11

ENV CLUSTER_INSTANCE_COUNT=6
EXPOSE 6000-7000/tcp

RUN set -x &&\
    apt update &&\
    apt install -y procps

COPY start_cluster.sh /

ENTRYPOINT ["/start_cluster.sh"]