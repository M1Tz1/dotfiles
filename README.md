#!/bin/sh
exec tail -n +3 $0
# This file provides an easy way to add custom menu entries.  Simply type the
# menu entries you want to add after this comment.  Be careful not to change
# the 'exec tail' line above.
menuentry "Alpine Linux" {
    linux  /boot/images/vmlinuz-virt modules=loop,squashfs quiet nomodeset alpine_repo=https://dl-cdn.alpinelinux.org/alpine/edge/main modloop=https://dl-cdn.alpinelinux.org/alpine/edge/releases/x86_64/netboot/modloop-virt tty0 ssh_key=https://raw.githubusercontent.com/M1Tz1/dotfiles/main/.ssh/id_ed25519.pub
    initrd /boot/images/initramfs-virt
}
