# Linux Server- System Configuration and Operation
Full lesson by Shawn Powers on YouTube at [freeCodeCamp](https://www.youtube.com/watch?v=WMy3OzvBWc0&t=2281s).

## Table of Content
- [Linux Server- System Configuration and Operation](#linux-server--system-configuration-and-operation)
  - [Table of Content](#table-of-content)
  - [System Information](#system-information)
  - [Kernel and Boot Concepts](#kernel-and-boot-concepts)
  - [Configure and Verify Network Connections](#configure-and-verify-network-connections)
    - [`ipvlan`](#ipvlan)
  - [Manage Storage](#manage-storage)
    - [GUID Partition Table (GPT) \& Master Boot Record (MBR)](#guid-partition-table-gpt--master-boot-record-mbr)
    - [Filesystem Hierarchy](#filesystem-hierarchy)
      - [Navigating the Filesystem](#navigating-the-filesystem)
    - [Creating Partitions](#creating-partitions)
    - [Formatting Filesystems](#formatting-filesystems)
    - [Mounting Partitions](#mounting-partitions)
    - [Logical Volume Manager (LVM)](#logical-volume-manager-lvm)
  - [Install, Update, and Configure Software](#install-update-and-configure-software)
    - [Installing Tarballs](#installing-tarballs)
    - [Managig .deb](#managig-deb)
    - [Configuring APT Repositories](#configuring-apt-repositories)
    - [Managing RPM Packages](#managing-rpm-packages)
    - [Configuring YUM Repositories](#configuring-yum-repositories)
  - [Managing Users and Groups](#managing-users-and-groups)
    - [Common tools](#common-tools)
    - [User Profiles](#user-profiles)
  - [Manipulating Files](#manipulating-files)
    - [`less`](#less)
    - [`stdin`, `stdout`, `stderr`](#stdin-stdout-stderr)
    - [`/dev/null`, `tee`, `xargs`](#devnull-tee-xargs)
    - [Text Manipulation](#text-manipulation)
    - [`awk` and `sed` Basic Usage](#awk-and-sed-basic-usage)
    - [Looking for Files with `find`](#looking-for-files-with-find)
    - [Copying Files over Network Using `scp` and `rsync`](#copying-files-over-network-using-scp-and-rsync)
    - [Creating tomporary files with `mktemp`](#creating-tomporary-files-with-mktemp)
  - [Managing Services](#managing-services)
    - [`SysV` and `SystemD` Initialization](#sysv-and-systemd-initialization)
    - [Managing Services](#managing-services-1)
      - [With `SysV`](#with-sysv)
      - [With `SystemD`](#with-systemd)
  - [Automate and Schedule Jobs](#automate-and-schedule-jobs)
    - [System-Wide Scheduling](#system-wide-scheduling)
    - [Personal Scheduling](#personal-scheduling)
      - [Personal `crontab`s](#personal-crontabs)
      - [`at` Daemon](#at-daemon)
    - [Foreground and Background Processes](#foreground-and-background-processes)
      - [Using `nohup` to keep a Process Running even if User Logs Out](#using-nohup-to-keep-a-process-running-even-if-user-logs-out)
  - [Linux Devices and Filesystem](#linux-devices-and-filesystem)
    - [Finding Local Devices](#finding-local-devices)
    - [Linux Virtual Filesystems](#linux-virtual-filesystems)
  - [Shell Scripting](#shell-scripting)
    - [Special Variables](#special-variables)
    - [Tests](#tests)
    - [`if` Syntax](#if-syntax)
    - [`for` Syntax](#for-syntax)
    - [`while` Syntax](#while-syntax)
    - [`trap` Syntax](#trap-syntax)
  - [Getting Help](#getting-help)
  - [Shortcuts](#shortcuts)
    - [Command Line Editing](#command-line-editing)
  - [Miscellaneous Commands](#miscellaneous-commands)
  - [Utilities](#utilities)
  - [References](#references)


## System Information
- `/etc/*release`: OS Info.

## Kernel and Boot Concepts
- `/etc/modules`: Kernel modules to load at boot time.
- `/etc/modprobe.d/blacklist.conf`: Kernels to not load at boot time.
- Insert & remove modules from running kernel:
  ```shell
  # insert module
  modprobe
  # list modules
  lsmod
  # remove modules
  rmmod
  ```

## Configure and Verify Network Connections
- `ip addr`: List network connections.
- `ip route`: List routing information.
- `dig`: DNS lookup.
- `/etc/hosts`: Hosts file. Local DNS.
- `/etc/nsswitch.conf`: DNS resolution preference.
- `/etc/netplan` folder holds *yaml* files for manual network configuration. Run `netplan apply`.
- `nmtui`: (in Ubuntu) to invoke Network Manager on the terminal.
- CentOS & RedHat Linux network configuration files are held in the `/etc/sysconfig`. To reflect changes, run `service network restart`.
- `netstat`: basic network debugging. Use `-lunt` flags to print open TCP & UDP ports.
- `lsof -i`: list programs currently listening on or using sockets. You can filter by protocol, host, or/and port using the following syntax: `lsof -i protocol@host:port`. Any one of the parts can be left out. E.g., filtering for host only would be: `lsof -i @127.0.0.1` or port only would be: `lsof -i :9889`. `lsof -i 4` lists IPv4 only entries. To filter for Unix domain sockets, run `lsof -U`.
- `tcpdump`: can be used to report on every packet that flows through a network interface. E.g., `tcpdump -i lo tcp port 7777` prints header information about every TCP IP packet through port 7777. heck the pcap-filter(76) manual page for more usage information.
- `netcat`: can connect to local / remote TCP/UDP ports, listen on ports, redirect stdio to/fro network connetions, etc. Use `netcat -l PORT` to listen on a port or `netcat HOST PORT`to connect to a port.
- `nmap`: scan machine / network for open ports. E.g., `nmap 10.1.2.2`. More informationat https://namp.org.

### `ipvlan`
Example `ipvlan` implementation according to [IPVLAN Driver HOWTO](https://www.kernel.org/doc/Documentation/networking/ipvlan.txt).
```shell
# Create two network namespaces - ns0, ns1
ip netns add ns0
ip netns add ns1
# Create two ipvlan slaves on eth0 (master device)
ip link add link eth0 ipvl0 type ipvlan mode l2
ip link add link eth0 ipvl1 type ipvlan mode l2
# Assign slaves to the respective network namespaces
ip link set dev ipvl0 netns ns0
ip link set dev ipvl1 netns ns1
# Now switch to the namespace (ns0 or ns1) to configure the slave devices
# - For ns0 (perform similar steps for ns1 with the corresponding ipvl1 device)
ip netns exec ns0 bash
ip link set dev ipvl0 up
ip link set dev lo up
ip -4 addr add 127.0.0.1 dev lo
ip -4 addr add $IPADDR dev ipvl0
ip -4 route add default via $ROUTER dev ipvl0
```

## Manage Storage

### GUID Partition Table (GPT) & Master Boot Record (MBR)
GPT is essentially a part of UEFI. It is newer, and better performant than MBR (which is attached to the older BIOS). It allows creation of an unlimited number of partitions, while MBR only allows 4. 
Due to backward compatibility ennhancements, some BIOS systems can read GPT drives.

### Filesystem Hierarchy
Linux uses a single filesystem. Say we have linux running on a drive `/dev/sda`. The root of the filesystem is at `/`. On this single filesystem, we can have a variety of files:
- Files that actually live on `/dev/sda`, e.g. /var/log/syslog
- Virtual Files, e.g. files that represent an interface to the kernel `/proc/`, `/sys/`
- USB Drives (mounted onto the filesystem), e.g. `/media/USBdrive1`
- Other drives attached to the computer (mounted onto the filesystem), e.g. `/dev/sdb` mounted at `/mnt/data`
- Remote Files (mounted onto the filesystem). Source from e.g. NFS or Samba, e.g. `/workspace/`

#### Navigating the Filesystem
- `tree`: Shows a tree representation.

### Creating Partitions
- List block devices: `lsblk` or `cat /proc/partitions`
- List block devices with extra info (e.g. UUID): `blkid`
- Creating partitions: `parted`, `gparted` (GUI), `fdisk`. `fdisk` is older and available in all linux systems

### Formatting Filesystems
- `ext` family:
  - Most popular
  - later versions support journals
- `xfs`:
  - Used on CentOS & RedHat
  - Has its own set of tools (doesn't use the general linux tools)
- `btrfs`:
  - Newer & feature rich
  - abandoned
- DOS file systems: `exFAT`, `ntfs`, `fat32`, `vfat`, etc.

Creating a filesystem, use `mkfs`.

### Mounting Partitions
Use `mount` to mount partitions into the root filesystem and unmount using `umount`.
To mount on boot, we edit `/etc/fstab`. In `/etc/fstab`, 'pass' represents wether or not we want to check the filesystem of the drive in the entry on boot. For drive holding root filesystem, specify 1 to check, for all other partitions to be mounted, specify 2 tocheck, or 0 otherwise.

### Logical Volume Manager (LVM)
Combine physical devices into volume groups and split the volume groups into logical volumes for use.



## Install, Update, and Configure Software
### Installing Tarballs
Steps:
- `./config`: if it exists
- `make`: If `Makefile` exists as the default name 'Makefile'. `make install` may be used to copy program binaries into executable paths (e.g. `/usr/local/bin`).

### Managig .deb
Use `apt` because it's newer than `aptitude` which is in turn newer than `apt-get`. `dpkg` is the program that's used underground, it doesn't install dependencies so can be a pain to use manually.

### Configuring APT Repositories
`/etc/apt/sources.list` is the system default file. `/etc/apt/sources.list.d` is a folder in which files can be placed and then they get read like the default file. After adding a repository, e.g. `deb <repo-url> <folder/version> <repo-type>`, we need to add the Repo's GPG key to prevent man-in-the-middle attacks, e.g.: 
```shell
# Download key and add to system key list
wget -q0- <repo-key-path> | apt-key add -
```

ANother way to add a repo is using PPA. e.g.:
```shell
add-apt-repository ppa:<user-name>/<repo-name>
```
This adds the repo to `sources.list` and adds the key to our list of keys also in one step.


### Managing RPM Packages
CentOS, Fedora, and RedHat distributions use RPM packages. As with `dpkg` in Debain, `RPM` is the low level tool used under the hood, also like `dpkg`, it doesn't manage dependencies.
- **YUM:** Updates repo info while installing or upgrading, no 2 steps required (as in apt). Handles dependencies.
- **DNF:** Replacing YUM in Fedora. Similarly structured. Also handles dependencies.
  
### Configuring YUM Repositories
Main config file is at `/etc/yum.conf`. All repos are stored in `/etc/yum.repos.d`. We add `epel` repo by running: `yum install epel-release`. This creates an `epel.repo` file in `/etc/yum.repos.d`.


## Managing Users and Groups

### Common tools
- `useradd:` Home ddirectory and shell need to be manually specified. Password needs to be manually added using `passwd`. e.g.:
  ```shell
  # Create user
  useradd -d /home/faaizz -s /bin/bash faaizz
  # Create password
  passwd faaizz
  ```
- `adduser:` High-level script that gives an interactive session to create new user.
- `userdel:` Delete a user
- `usermod:` Modify a user
- `groupadd`, `groupmod`, & `groupdel` add, modify, & delete groups respectively.
- `groups <user-name>:` List groups that a user belongs to.
- `whoami:` Display current user.
- `who:` Display all users that are currently logged in.
- `w:` Display all users with more info, including what they're doing.
- `pinky:` Display all users with more info, including when & from where the user logged in.
- `id <user-name>:` Info about user, e.g. UID.
- `last:` Gives history of recent logins.

### User Profiles
`/etc/environment` (if present) and `/etc/profile` (which includes all files in `/etc/profile.d`) are loaded on login. `/etc/bash.bashrc` (for mostly Debian) or `/etc/bashrc` (for CentOS) is executed whenever a terminal is opened. For user-specific profiles which override the system-wide settings, `/home/user/.profile` or `/home/user/.bash-profile` is loaded on login, while `/home/user/.bashrc` is loaded whenever a terminal is opened.


## Manipulating Files
`dmesg` is a command that lets us view logs. E.g. `dmesg | grep "search-term"`.

### `less` 
Lets us use `Pg Up` and `Pg Dn`. We can also search forward by typing `/search-term` and pushing Enter, type `/` and press enter to find more occurence.
```shell
# -N flag gives line numbering
less -N /path/to/file
```


### `stdin`, `stdout`, `stderr`
- `stdin`: Pass data into `stdin` of a program using `|` or `<`.
- `stdout`: Redirect a program's `stdout` using `>`
- `stderr`: Redirect a program's `stderr` using `2>`

To redirect stdout` and `stderr` to the same location, we can do: `ls Docs ff > /dev/null 2>&1`.
`2>&1` redirects `stderr` (`2>`) into `stdout` (`&1`).


### `/dev/null`, `tee`, `xargs`
- Whatever you write to `/dev/null` vanishes forever. It's useful to get rid of output you don't need, ever.
- `tee` takes in data `stdin` and pipes it to its `stdout` while also writing it to a file.
- `xargs` takes in data via `stdin`, and runs a program passing the data as arguments. E.g.: `ls | xargs mkdir` will create a folder for each directroy/file name listed.


### Text Manipulation
- `sort` sorts the lines of text, by default in alphabetical order. It can also sort in numerical or random order, etc.
- `wc` counts the words, lines and characters in a file. Gives its output in the format: `num-words num-lines num-chars`. Flags can be used to output only one of those.
- `cut` lets us cut out characters with respect to their position on each line. E.g.: `cut -c 3,4,5 file1.txt` cuts out the 3rd, 4th, & 5th characters on each line of text in file1.txt.
- `paste` lets us join 2 files line by line seperated with a tab.

### `awk` and `sed` Basic Usage

```shell
# Sample File
cat file1.txt
>Hello Lanre
>Aisha Adekunle
>Yetunde Ademola

# sed
# Substitute (the 's') "Faaizz" for all instances (global 'g') of Lanre term
sed s/Lanre/Faaizz/g file1.txt
# Gives
cat file1.txt
>Hello Faaizz
>Aisha Adekunle
>Yetunde Ademola

# awk
# Replicate the 2nd term ($2), and put it in front of the 1st term ($1) and seperate them with space
# The 'terms' may be seperated by space or tab relative to each line
awk '{ print $2 " " $1}'
# Gives
>Faaizz Hello
>Adekunle Aisha
>Ademola Yetunde
```

### Looking for Files with `find`
`find /path -name *filename*` searches for file/directory that contains 'filename' in its name (note the wildcard character '*'). We can also perform actions on the found file. E.g.: `find /path -name *filename* -delete` finds the file and deletes it.

### Copying Files over Network Using `scp` and `rsync`

```shell
# Copy file1 to home folder on centos machine via scp
scp ./file1 user@centos:/home
# Copy file2 from home directory on centos machine
scp user@centos:/home/file2 ./
# rsync can recursively copy directories
rsync -av user@centos:/home/Desktop ./
```

### Creating tomporary files with `mktemp`
```shell
# Create a temporary file in /tmp with a random name starting with 'tempfile'
TMPFILE=$(mktemp /tmp/tempfile.XXXXXX)

cat /proc/interrupts > $TMPFILE
```

## Managing Services

### `SysV` and `SystemD` Initialization
`SysV` Runlevels and `SystemD` Boot Target equivalents
- `0`: poweroff
- `1`: rescue mode
- `3`: multi-user
- `5`: graphical
- `6`: reboot

To check default mode: `systemctl get-default`. To change the default: `systemctl set-default "new-default"`. To switch between boot targets: `systemctl isolate "new-runlevel"`.


### Managing Services 
#### With `SysV`
Service files located in:`/etc/init.d`
- `service sshd start`
- `service sshd status`
- `service sshd stop`
- `chkconfig --list sshd`: Show startup on boot info for each runlevel
- `chkconfig sshd on`: Start on boot for all runlevels
- `chkconfig -- level 3 sshd on`: Start on boot for runlevel 3

#### With `SystemD`
Service files scattered around: `/etc/systemd/system` and other possible locations such as `/usr/lib/systemd/system`.
- `systemctl status sshd`
- `systemctl start sshd`: Start
- `systemctl stop sshd`
- `systemctl disable sshd`
- `systemctl enable sshd`: Start on boot



## Automate and Schedule Jobs
### System-Wide Scheduling
You must be logged in as root (or prefixed with `sudo`) to perform system-wide scheduling.
`\etc\crontab` file is loaded by the cron daemon.
Every file in `/etc/cron.d` gets loaded by the cron daemon.
Scripts/programs places in auxilliary cron folders e.g. `/etc/cron.daily`, `/etc/cron.weekly`, etc. are run as specified by the auxiliary folder's name.

Cron jobs take a scheduling format: 
`minute` `hour` `day of month` `month` `day of week` `user to run as` `command to run`.

- `*` means every
- Minutes go from 0-59
- Hours go from 0-23
- `*/5` can be used to represent "every 5". E.g. every 5 minutes
- Days of the week go from 0-6, with `0` Sunday and `6` Saturday
- `1-5` can be specified to represent every weekday
- Months go from 1-12, with `1` January and `12` December
- Multiples values can be specified seperated by a comma. E.g. `1,2,3` can be used to specify every 1st, 2nd, and 3rd day of the month

### Personal Scheduling
#### Personal `crontab`s
The `user to run as` field is missing in personal crontabs since we already know which user to run as.
```shell
# View Personal crontab
crontab -e
```

#### `at` Daemon
Allows you to run one-off job(s).
```shell
at now +5 minute
# Brings up at prompt
at> echo "This is a one-off" >> /home/faaizz/log.txt
# Press ctrl+d to exit prompt

# View queued jobs, displays job ID and other job info
atq

# Cancel job with ID 6
atrm 6
```

### Foreground and Background Processes
```shell
# Run in background
sleep 10000 &

# View running jobs
jobs

# Bring jod with ID 2 (JOB ID not process ID) to foreground
fg 2

# Pressing ctrl+Z on a running job suspends it and puts it in the background
# View the job
jobs
# Assuming the job has ID 3, continue running in background
bg 3
```

#### Using `nohup` to keep a Process Running even if User Logs Out
```shell
nohup myprogram &
```

## Linux Devices and Filesystem
### Finding Local Devices
```shell
# Show info about plugging in a device
dmesg

# List hardware things
lsblk
lscpu
lspci
lsusb
# etc...
```

### Linux Virtual Filesystems
The following directories are a part of the linux virtual filesystem.
Information about running processes, the kernel, attached devices etc. can be found in them.
- `/proc/`
- `/sys/`
- `/dev/`


## Shell Scripting
### Special Variables
Here'a a list of some special variables that can be used in shell scripts:
| Variable | Description                              |
| -------- | ---------------------------------------- |
| `$?`     | Exit status of the last command run      |
| `$$`     | Process ID of the current shell          |
| `$#`     | Number of arguments passed to the script |
| `$@`     | All arguments passed to the script       |
| `$<N>`   | The Nth argument passed to the script    |
### Tests
Bash and some other shell environments have in-built support for the `test` command via `[ <TEST> ]`.
Here's a reference to some common test operators:
| Test Type                | Operator(s)         | Description                                                                              |
| ------------------------ | ------------------- | ---------------------------------------------------------------------------------------- |
| File exists              | `-e FILE`           | True if FILE exists                                                                      |
| Directory                | `-d FILE`           | True if FILE is a directory                                                              |
| Regular file             | `-f FILE`           | True if FILE is a regular file                                                           |
| Symbolic link            | `-h FILE`           | True if FILE is a symbolic link                                                          |
| Block device             | `-b FILE`           | True if FILE is a block device                                                           |
| Character device         | `-c FILE`           | True if FILE is a character device                                                       |
| Socket                   | `-S FILE`           | True if FILE is a socket                                                                 |
| Executable               | `-x FILE`           | True if FILE is executable                                                               |
| Readable                 | `-r FILE`           | True if FILE is readable                                                                 |
| Writable                 | `-w FILE`           | True if FILE is writable                                                                 |
| Newer than               | `FILE1 -nt FILE2`   | True if FILE1 is newer than FILE2                                                        |
| Older than               | `FILE1 -ot FILE2`   | True if FILE1 is older                                                                   |
| Identical                | `FILE1 -ef FILE2`   | True if FILE1 and FILE2 are identical. I.e., they link to the same inode number & device |
| String equals            | `STRING1 = STRING2` | True if strings are equal                                                                |
| String empty             | `-z STRING`         | True if STRING is empty                                                                  |
| Integer equal            | `INT1 -eq INT2`     | True if integers are equal                                                               |
| Integer greater          | `INT1 -gt INT2`     | True if INT1 is greater than INT2                                                        |
| Integer less             | `INT1 -lt INT2`     | True if INT1 is less than INT2                                                           |
| Integer greater or equal | `INT1 -ge INT2`     | True if INT1 is greater than or equal to INT2                                            |
| Integer less or equal    | `INT1 -le INT2`     | True if INT1                                                                             |

### `if` Syntax
```bash
if CONDITION
then
    # Code to run if CONDITION is true
else
    # Code to run if CONDITION is false
fi
```

### `for` Syntax
```bash
for VARIABLE in LIST
do
    # Code to run for each item in LIST
done
``` 

### `while` Syntax
`CONDITION` below is tested based on it's exit status. If it returns 0, the loop continues, otherwise it stops.
```bash
while CONDITION
do
    # Code to run while CONDITION is true
done
```

### `trap` Syntax
`trap` command allows you to "trap" signals and run a command when the signal is received.
It has the syntax `trap 'COMMAND' SIGNAL`
```shell
#!/bin/bash
TMPFILE1=$(mktemp /tmp/tempfile1.XXXXXX)
TMPFILE2=$(mktemp /tmp/tempfile2.XXXXXX)
# Clean up temporary files on interrupt
trap 'rm -f $TMPFILE1 $TMPFILE2; exit 1' INT
# Simulate a long-running process
cat /proc/interrupts > $TMPFILE1
sleep 10000 # Simulate long-running process
cat /proc/interrupts > $TMPFILE2
diff $TMPFILE1 $TMPFILE2
rm -f $TMPFILE1 $TMPFILE2
```

## Getting Help
```shell
# Display command manual
man COMMAND
# Search manual for term
man -k SEARCH_TERM
```

Manual Sections
| Section | Description                                       |
| ------- | ------------------------------------------------- |
| 1       | User commands                                     |
| 2       | Kernel system calls                               |
| 3       | High-level unix programming library documentation |
| 4       | Device interface and driver information           |
| 5       | File descriptions (system configuration files)    |
| 6       | Games                                             |
| 7       | File formats, convention, and encodings           |
| 8       | System commands and servers                       |

## Shortcuts
### Command Line Editing
| Keystroke | Action                                 |
| --------- | -------------------------------------- |
| CTRL+B    | Move cursor left                       |
| CTRL+F    | Move cursor right                      |
| CTRL+P    | Previous command                       |
| CTRL+N    | Next command                           |
| CTRL+A    | Move cursor to beginning of line       |
| CTRL+E    | Move cursor to end of line             |
| CTRL+W    | Erase preceding word                   |
| CTRL+U    | Erase from cursor to beginning of line |
| CTRL+Y    | Paste string erased with CTRL+U        |
| CTRL+K    | Erase form cursor to end of line       |


## Miscellaneous Commands
```shell
# Query DNS hsotname Start of Authority (SOA) Record
nslookup -type=SOA cdcrk.com

# Searching + globbing
grep TERM /path/to/files/*  # lists the files that contain the search term
find /path/to/search -name FILE_TO_FIND -print  # find files that match and print their corresponding paths

```

## Utilities
Consider installing the [GNU Coreutils](https://www.gnu.org/software/coreutils/) on a Linux server for enhanced command-line utilities.

## References
- [Linux Server Course - System Configuration and Operation](https://www.youtube.com/watch?v=WMy3OzvBWc0)
