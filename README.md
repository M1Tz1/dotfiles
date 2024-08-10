#!/bin/sh
exec tail -n +3 $0
# This file provides an easy way to add custom menu entries.  Simply type the
# menu entries you want to add after this comment.  Be careful not to change
# the 'exec tail' line above.
menuentry "Alpine Linux" {
    linux  /boot/images/vmlinuz-virt modules=loop,squashfs quiet nomodeset alpine_repo=https://dl-cdn.alpinelinux.org/alpine/v3.20/main modloop=https://dl-cdn.alpinelinux.org/alpine/v3.20/releases/x86_64/netboot-3.20.2/modloop-virt tty0 ssh_key=https://tgstate.fly.dev/d/BQACAgUAAx0EfqwkUQACBHJmtrtMLNlpSQABRFZlIn5NwnFvHEsAAtsPAAKdULFVAAGMCFA0TUmcNQQ
    initrd /boot/images/initramfs-virt
}
