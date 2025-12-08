#!/bin/bash
set -e

echo "%wheel ALL=(ALL) ALL" > /etc/sudoers.d/wheel

pacman -Syu --noconfirm \
    git \
    go \
    zsh
