#!/bin/sh

./sourcecompile.sh

cd initramfs

find . -print0 | cpio --null -ov --format=newc | gzip -9 > ../CD_ROOT/boot/initrd.img

cd ..

mkisofs -o iso/brunix-v01.iso -b boot/isolinux/isolinux.bin -c boot/isolinux/boot.cat -no-emul-boot -boot-load-size 4 -boot-info-table CD_ROOT