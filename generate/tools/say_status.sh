#!/bin/bash

PROJECT="darwinkit"
TASK="documentation update"
VOICE="Samantha"

# Create directories
mkdir -p /tmp/.composer-say

# Record voice choice
echo "${PROJECT}:${TASK}:${VOICE}" >> /tmp/.composer-project-voices

# Try to acquire lock
if [ ! -f /tmp/.composer-say/lock ]; then
    touch /tmp/.composer-say/lock
    say -v "${VOICE}" "${PROJECT} ${TASK}: $1"
    sleep 1
    rm /tmp/.composer-say/lock
fi