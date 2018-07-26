#!/usr/bin/env bash

set -eu -o pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)

DEPLOYMENT_NAME="${DEPLOYMENT_NAME:="ci-service"}"

export GOPATH="${ROOT}/git-kubo-ci"

main() {
  BOSH_ENVIRONMENT="$(jq '.target' "${ROOT}/gcs-source-json/source.json")"
  BOSH_CLIENT="$(jq '.client' "${ROOT}/gcs-source-json/source.json")"
  BOSH_CLIENT_SECRET="$(jq '.client_secret' "${ROOT}/gcs-source-json/source.json")"
  BOSH_CA_CERT="$(jq '.ca_cert' "${ROOT}/gcs-source-json/source.json")"
  BOSH_DEPLOYMENT="$DEPLOYMENT_NAME"

  export BOSH_ENVIRONMENT BOSH_CLIENT BOSH_CLIENT_SECRET BOSH_CA_CERT BOSH_DEPLOYMENT

  CONFIG="${ROOT}/gcs-kubeconfig/k8s/config" ginkgo -r -progress "${ROOT}/git-kubo-ci/src/tests/bbr-tests/"
}

main
