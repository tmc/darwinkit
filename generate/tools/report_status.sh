#!/bin/bash
mkdir -p /tmp/compose-say
if [ ! -f "/tmp/compose-say/darwinkit.lock" ]; then touch "/tmp/compose-say/darwinkit.lock"; say "$1"; sleep 2; rm "/tmp/compose-say/darwinkit.lock"; fi
