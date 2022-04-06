%define debug_package %{nil}

Name:     hello
Version:  %{_version}
Release:  %{_release}%{?dist}
Summary:  hello
License:  Apache 2.0
URL:      https://github.com/fengye87/hello
Source0:  hello
Source1:  hello.service
Source2:  hello.preset

%description
hello

%install
install -d -m 755 %{buildroot}%{_bindir}
install -c -m 755 %{SOURCE0} %{buildroot}%{_bindir}/hello

install -d -m 755 %{buildroot}%{_unitdir}
install -c -m 644 %{SOURCE1} %{buildroot}%{_unitdir}/hello.service

install -d -m 755 %{buildroot}%{_presetdir}
install -c -m 644 %{SOURCE2} %{buildroot}%{_presetdir}/hello.preset

%files
%defattr(-,root,root,-)
%{_bindir}/hello
%{_unitdir}/hello.service
%{_presetdir}/hello.preset

%changelog
