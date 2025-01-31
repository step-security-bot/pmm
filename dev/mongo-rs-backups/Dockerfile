FROM percona/percona-server-mongodb:4.2@sha256:ff03389831a8ec3567389732cef8e45bb4fa1ac6e972ee4a2a5e0218fd6da1ce

USER root

RUN yum install -y https://repo.percona.com/yum/percona-release-latest.noarch.rpm && \
    percona-release enable pbm release && \
    percona-release enable pmm2-client release

RUN yum install -y percona-backup-mongodb pmm2-client

