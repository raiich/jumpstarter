#!/bin/bash
set -euo pipefail

VERSION="${VERSION:-latest}"

# Provided by the devcontainer CLI (see containers.dev/implementors/features)
USERNAME="${_REMOTE_USER:-vscode}"
USER_HOME="${_REMOTE_USER_HOME:-/home/${USERNAME}}"

# Download and run the native installer as the container user
su - "${USERNAME}" -c "curl -fsSL https://claude.ai/install.sh | bash -s ${VERSION}"
