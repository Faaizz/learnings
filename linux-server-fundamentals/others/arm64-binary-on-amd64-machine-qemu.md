# Run an ARM64 Binary on an AMD64 Machine using QEMU Emulation
The Linux Kernel supports configuring custom interpreters for arbritrary binary formats via the [binfmt_misc](https://docs.kernel.org/admin-guide/binfmt-misc.html) feature.
[QEMU User Space Emulation](https://www.qemu.org/docs/master/user/main.html) can be used to run individual binaries that were compiled for foreign architectures.
Registering the QEMU User Space Emulator for `aarch64` binaries via binfmt_misc means the kernel would intercept any attempt to execute ARM64 binaries and route it to the emulator.
### Sample
```shell
# Install via APT & registers custom interpreters via binfmt_misc
sudo apt install -y qemu-user-binfmt
# Check registered interpreters
ls /proc/sys/fs/binfmt_misc
# Try executing any ARM64 binary
# ...
# It would be intercepted and /usr/bin/qemu-aarch64 would be run instead
#   with the biinary as an argument
# Find the process PID and check process executable with
ls -l /proc/<PID>/exe
```