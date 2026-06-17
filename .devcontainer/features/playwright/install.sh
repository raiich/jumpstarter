#!/bin/bash
set -euo pipefail

VERSION="${VERSION:-latest}"
BROWSERS="${BROWSERS:-chromium}"

# Provided by the devcontainer CLI (see containers.dev/implementors/features)
USERNAME="${_REMOTE_USER:-vscode}"

export DEBIAN_FRONTEND=noninteractive
# Shared path (exported to the container via containerEnv) so browsers installed
# here as root at build time are found by the container user at runtime.
export PLAYWRIGHT_BROWSERS_PATH=/usr/local/share/ms-playwright

# `playwright install` registers the installing package in the cache's .links,
# and later installs garbage-collect builds whose registered package is gone.
# Install the package persistently — not via the npx cache — so the baked
# browsers keep a live registration: project-level installs of other versions
# then add their builds alongside instead of deleting these.
# Requires node on PATH: the node feature must install before this one
# (installsAfter / overrideFeatureInstallOrder).
CLI_DIR=/usr/local/share/playwright-cli
npm install --prefix "${CLI_DIR}" "playwright@${VERSION}"

# shellcheck disable=SC2086  # BROWSERS is a space-separated list
"${CLI_DIR}/node_modules/.bin/playwright" install --with-deps ${BROWSERS}

# Owned by the container user so a project can add browsers or versions later.
chown -R "${USERNAME}:" "${PLAYWRIGHT_BROWSERS_PATH}"

# --with-deps runs apt-get; npm leaves a download cache under /root.
rm -rf /var/lib/apt/lists/* /root/.npm
