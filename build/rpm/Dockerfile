FROM fedora:35

RUN dnf install -y fedora-packager && dnf clean all

RUN case $(uname -m) in x86_64) arch=amd64;; aarch64) arch=arm64;; esac; \
    curl -SL https://golang.org/dl/go1.17.8.linux-$arch.tar.gz | tar -xzC /usr/local
ENV PATH /usr/local/go/bin:$PATH
