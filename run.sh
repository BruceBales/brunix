#!/bin/bash

./sourcecompile.sh

cd initramfs-new

find . -print0 | cpio --null -ov --format=newc | gzip -9 > ../initramfs.cpio.gz

cd ..

sudo qemu-system-x86_64 --enable-kvm -kernel ./linux/bzImage --initrd initramfs.cpio.gz -m 2G