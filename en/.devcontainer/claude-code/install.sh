#!/bin/bash
set -euo pipefail

VERSION="${VERSION:-latest}"

# Provided by the devcontainer CLI (see containers.dev/implementors/features)
USERNAME="${_REMOTE_USER:-vscode}"
USER_HOME="${_REMOTE_USER_HOME:-/home/${USERNAME}}"

# Download and run the native installer as the container user
su - "${USERNAME}" -c "curl -fsSL https://claude.ai/install.sh | bash -s ${VERSION}"

# Ensure .local/bin is in PATH (only if not already present)
for shell_rc in ".bashrc" ".profile" ".zshrc"; do
    rc_path="${USER_HOME}/${shell_rc}"
    if [ -f "${rc_path}" ] && ! grep -q '\.local/bin' "${rc_path}"; then
        printf '\nexport PATH="${HOME}/.local/bin:${PATH}"\n' >> "${rc_path}"
    fi
done
