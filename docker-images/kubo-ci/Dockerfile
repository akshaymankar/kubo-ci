FROM ubuntu:16.04
MAINTAINER pcf-kubo@pivotal.io

# Packages
RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install -y \
  curl \
  gcc \
  jq \
  git-core \
  netcat-openbsd  \
  make \
  python-pip \
  python-dev \
  python-software-properties \
  software-properties-common \
  wget \
  ipcalc \
  zip \
  vim \
  haproxy \
  iptables # required for sshuttle

WORKDIR /tmp/docker-build

# Golang
ENV GOLANG_VERSION=1.10.3
RUN wget https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV GOPATH /root/go
RUN mkdir -p /root/go/bin
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
RUN go get github.com/onsi/gomega && \
   go get github.com/onsi/ginkgo/ginkgo && \
   go get github.com/tsenart/vegeta

# CF
RUN wget -O /cf-cli_amd64.deb https://cli.run.pivotal.io/stable?release=debian64 && \
 dpkg -i /cf-cli_amd64.deb && rm /cf-cli_amd64.deb

# om

RUN wget -q -O - https://raw.githubusercontent.com/starkandwayne/homebrew-cf/master/public.key | apt-key add - && \
echo "deb http://apt.starkandwayne.com stable main" | tee /etc/apt/sources.list.d/starkandwayne.list &&\
apt-get update && apt-get install om

# Google SDK
ENV GCLOUD_VERSION=211.0.0
ENV GCLOUD_SHA1SUM=53becbf0a512bc1f928921a116c9a397cbf172621849e1d85dd2f906933b2241

RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-${GCLOUD_VERSION}-linux-x86_64.tar.gz \
    -O gcloud_${GCLOUD_VERSION}_linux_amd64.tar.gz && \
    echo "${GCLOUD_SHA1SUM}  gcloud_${GCLOUD_VERSION}_linux_amd64.tar.gz" > gcloud_${GCLOUD_VERSION}_SHA1SUM && \
    shasum -a 256 -cw --status gcloud_${GCLOUD_VERSION}_SHA1SUM && \
    tar xvf gcloud_${GCLOUD_VERSION}_linux_amd64.tar.gz && \
    mv google-cloud-sdk / && cd /google-cloud-sdk  && ./install.sh

ENV PATH=$PATH:/google-cloud-sdk/bin

# Ruby required for bosh-cli create-env
RUN git clone https://github.com/postmodern/ruby-install.git /usr/local/ruby-install
ENV PATH /usr/local/ruby-install/bin:/usr/local/bin:$PATH
RUN ruby-install --system ruby 2.5.1

# Credhub
RUN wget https://github.com/cloudfoundry-incubator/credhub-cli/releases/download/1.7.7/credhub-linux-1.7.7.tgz \
  -O credhub-linux-1.7.7.tgz && tar xvf credhub-linux-1.7.7.tgz && mv credhub /usr/bin

# kubectl
ENV KUBE_VERSION="v1.11.2"
ADD https://storage.googleapis.com/kubernetes-release/release/$KUBE_VERSION/bin/linux/amd64/kubectl /usr/bin/kubectl
RUN chmod +x /usr/bin/kubectl

# BOSH CLI
RUN wget https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-5.1.1-linux-amd64 -O bosh && \
  echo eccaaf1f5d9e0399271403c6966108443367ac80 bosh > bosh-shasum && \
  sha1sum -cw --status bosh-shasum && \
  install bosh /usr/local/bin

# Openstack CLI
RUN pip install -U setuptools
RUN pip install cryptography==2.0.3
RUN pip install pyOpenSSL==17.3.0
RUN pip install python-glanceclient==2.8.0
RUN pip install python-openstackclient==3.13.0

# AWS CLI
RUN pip install awscli

# sshuttle
RUN pip install sshuttle

# GOVC CLI
RUN wget -O - -o /dev/null https://github.com/vmware/govmomi/releases/download/v0.17.1/govc_linux_amd64.gz | gunzip > /usr/local/bin/govc && \
  chmod +x /usr/local/bin/govc

RUN gem install bundler --no-ri --no-rdoc

ARG SPRUCE_VERSION=v1.16.2
RUN wget https://github.com/geofffranks/spruce/releases/download/${SPRUCE_VERSION}/spruce-linux-amd64 -O /usr/bin/spruce && \
  chmod +x /usr/bin/spruce

RUN git clone https://github.com/fsaintjacques/semver-tool && \
  cd semver-tool && make install

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
