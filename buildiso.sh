#!/bin/sh

./sourcecompile.sh

cd initramfs-busybox

find . -print0 | cpio --null -ov --format=newc | gzip -9 > ../initrd.img

cd ..

mkisofs -o iso/brunix.iso -b boot/isolinux/isolinux.bin -c boot/isolinux/boot.cat -no-emul-boot -boot-load-size 4 -boot-info-table CD_ROOT