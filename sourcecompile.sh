go build -o ./initramfs-new/init ./src/init/cmd

go build -o ./initramfs-new/bin/mount ./src/mount/cmd

go build -o ./initramfs-new/bin/gsh ./src/gsh/cmd
