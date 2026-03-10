#!/bin/bash
set -euo pipefail

# Install firewall packages
apt-get update && apt-get install -y --no-install-recommends \
    iptables \
    ipset \
    iproute2 \
    dnsutils \
    jq \
    && rm -rf /var/lib/apt/lists/*

# Deploy init-firewall.sh
cp "$(dirname "$0")/init-firewall.sh" /usr/local/bin/init-firewall.sh
chmod +x /usr/local/bin/init-firewall.sh

# Allow the remote user to run init-firewall.sh via sudo without password
USERNAME="${_REMOTE_USER:-vscode}"
echo "${USERNAME} ALL=(root) NOPASSWD:/usr/local/bin/init-firewall.sh" \
    > /etc/sudoers.d/init-firewall
chmod 0440 /etc/sudoers.d/init-firewall
