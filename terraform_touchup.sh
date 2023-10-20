#!/usr/bin/env bash
set -xe

# This script needs to be run in order to populate the list of environments for each account in atlantis.yaml.
# We've resorted to using this script as a workaround to having to do terraform workspace list in Github Actions

TERRAFORM_VERSION="1.3.4"

CURRENT_DIR=$(pwd)

rm atlantis.yaml
touch atlantis.yaml

{
  echo "version: 3"
  echo "automerge: false"
  echo "projects:"
} >> atlantis.yaml

# RUN THIS SCRIPT BEFORE PUSHING THE PULL REQUEST
if [[ -d "deploy/tf" ]]; then
  for dir in deploy/tf/accounts/*; do
    unset envs
    cd "$CURRENT_DIR/$dir"
    ACCOUNT=$(basename "${dir}")
    terraform init
    terraform providers lock -platform=windows_amd64 -platform=darwin_amd64 -platform=linux_amd64
    git stage .terraform.lock.hcl
    pwd
    envs=($(terraform workspace list))
    echo "${envs[@]}"
    for env in "${envs[@]}"
    do
      :
      if [[ $env != "*" && $env != "default" && $env != *".tf"* ]]; then
        cd "$CURRENT_DIR"
        {
          echo "  - name: tf-${ACCOUNT}-${env}"
          echo "    dir: ./deploy/tf/accounts/${ACCOUNT}"
          echo "    workspace: ${env}"
          echo "    terraform_version: 1.3.4"
        } >> atlantis.yaml
        cd "$CURRENT_DIR/$dir"
      fi
    done
  done
  cd "$CURRENT_DIR"
fi

if [[ -d "deploy/tf_application" ]]; then
  for dir in deploy/tf_application/accounts/*; do
    unset envs
    cd "$CURRENT_DIR/$dir"
    ACCOUNT=$(basename "${dir}")
    terraform init
    terraform providers lock -platform=windows_amd64 -platform=darwin_amd64 -platform=linux_amd64
    git stage .terraform.lock.hcl
    pwd
    envs=($(terraform workspace list))
    echo "${envs[@]}"
    for env in "${envs[@]}"
    do
      :
      if [[ $env != "*" && $env != "default" && $env != *".tf"* ]]; then
        cd "$CURRENT_DIR"
        {
          echo "  - name: tf-${ACCOUNT}-${env}"
          echo "    dir: ./deploy/tf_application/accounts/${ACCOUNT}"
          echo "    workspace: ${env}"
          echo "    terraform_version: 1.3.4"
        } >> atlantis.yaml
        cd "$CURRENT_DIR/$dir"
      fi
    done
  done
  cd "$CURRENT_DIR"
fi

terraform fmt --recursive

# Auto delete the script after successful use
rm terraform_touchup.sh
