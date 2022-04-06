#!/bin/bash
set -euxo pipefail

version=$(git describe --long --tags --dirty | awk '{print substr($1,2)}')
rpm_version=$(echo $version | sed -nr "s/^([0-9]+(\.[0-9]+)+)(-([-A-Za-z0-9\.]+))?$/\1/p")
rpm_release=$(echo $version | sed -nr "s/^([0-9]+(\.[0-9]+)+)(-([-A-Za-z0-9\.]+))?$/\4/p" | sed 's/-/./g')

mkdir -p /root/rpmbuild/SOURCES
rm -rf /root/rpmbuild/SOURCES/*

cp build/rpm/hello.service /root/rpmbuild/SOURCES
cp build/rpm/hello.preset /root/rpmbuild/SOURCES

for name in hello helloctl; do
    go build -o /root/rpmbuild/SOURCES/$name cmd/$name/main.go
done

mkdir -p out/rpm

for name in hello helloctl; do
    cp build/rpm/$name.spec /root/rpmbuild/SOURCES/$name.spec
    git log --format="* %cd %aN%n- (%h) %s%d%n" -n 10 --date local | sed -r 's/[0-9]+:[0-9]+:[0-9]+ //' >> /root/rpmbuild/SOURCES/$name.spec
    rpmbuild --define "_version $rpm_version" --define "_release $rpm_release" -bb /root/rpmbuild/SOURCES/$name.spec
    cp /root/rpmbuild/RPMS/**/$name-*.rpm out/rpm
done
