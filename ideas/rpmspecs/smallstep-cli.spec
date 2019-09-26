Name:           smallstep-cli
Version:        0.13 
Release:        2
Summary:        A zero trust swiss army knife for working with X509, OAuth, JWT, OATH OTP, etc. 

Group:          Security
BuildArch:      x86_64 
License:        Apache License, Version 2.0 
URL:            https://github.com/smallstep/cli 
Source0:        smallstep-cli-0.13.2.tar.gz 

%description
step is a zero trust swiss army knife that integrates with step-ca for automated certificate management. It's an easy-to-use and hard-to-misuse utility for building, operating, and automating systems that use zero trust technologies like authenticated encryption (X.509, TLS), single sign-on (OAuth OIDC, SAML), multi-factor authentication (OATH OTP, FIDO U2F), encryption mechanisms (JSON Web Encryption, NaCl), and verifiable claims (JWT, SAML assertions).

# The idea is to get to where the rpm will build from source or get the binarys from the smallstep repos
%prep
%setup -q
%global _missing_build_ids_terminate_build 0
%install
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/bin
install -m 0644 README.md $RPM_BUILD_ROOT/usr/local/bin/SMALLSTEP-CLI-README.md
install -m 0755 step $RPM_BUILD_ROOT/usr/local/bin/step

%files
/usr/local/bin/SMALLSTEP-CLI-README.md
/usr/local/bin/step

%changelog
* Thu Sep 26 2019 Inital packaging Matthew Feinberg (matthew.feinberg@rmgmedia.com) 
  - Initial rpm release
