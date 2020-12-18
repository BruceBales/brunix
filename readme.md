
Compile main.go into initramfs-go/init

Commands to run:

Copy `initramfs-go` folder to a usable initramfs image:
```
cd initramfs-busybox

find . -print0 | cpio --null -ov --format=newc | gzip -9 > ../initramfs.cpio.gz

cd ..

```

Run the OS in qemu:
```
qemu-system-x86_64 --enable-kvm -kernel ./linux/bzImage --initrd initramfs.cpio.gz
```