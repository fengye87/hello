%define debug_package %{nil}

Name:     helloctl
Version:  %{_version}
Release:  %{_release}%{?dist}
Summary:  helloctl
License:  Apache 2.0
URL:      https://github.com/fengye87/hello
Source0:  helloctl

%description
helloctl

%install
install -d -m 755 %{buildroot}%{_bindir}
install -c -m 755 %{SOURCE0} %{buildroot}%{_bindir}/helloctl

%files
%defattr(-,root,root,-)
%{_bindir}/helloctl

%changelog
