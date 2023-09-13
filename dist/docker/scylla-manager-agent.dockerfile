FROM --platform=linux/arm64 ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY dist/release/scylla-manager-agent*.linux_arm64.deb /
RUN dpkg -i scylla-manager-agent*.deb && rm /scylla-manager-agent*.deb

USER scylla-manager
ENV HOME /var/lib/scylla-manager/
ENTRYPOINT ["/usr/bin/scylla-manager-agent"]
