# LPIC-2 Linux Engineer Cert Prep

Notes from LPIC-2 Linux Engineer (201-450) Cert Prep video course on LinkedIn Learning.
*Note: A debian-based linux distro is assumed for the listed commands.*

## Content
- [LPIC-2 Linux Engineer Cert Prep](#lpic-2-linux-engineer-cert-prep)
  - [Content](#content)
  - [Supporting Linux](#supporting-linux)
    - [Capacity Planning](#capacity-planning)
      - [Measuring CPU Activity](#measuring-cpu-activity)
      - [Measuring Memory Usage](#measuring-memory-usage)
      - [Measuring Disk Activity](#measuring-disk-activity)

## Supporting Linux
### Capacity Planning
4 main metrics:
- CPU
- Memory (RAM)
- Disk I/O
- Network

Relevant Tools:
```shell
# Monitor real-time consumption
## Ships with most linux distros
top
## More fancy/graphical
htop

# Review historical consumption: sar
## Install sysstat package
apt install sysstat
## Enable logging: set ENABLED="true"
sudoedit /etc/default/sysstat
## Adjust logging frequency as required. Data is stored in /var/log/sysstat
sudoedit /etc/cron.d/sysstat
## View logged data
### CPU usage
sar -u
### Memory usage
sar -r
### Disk I/O
sar -b
### Network
sar -n
### Specific time frame
#### <file_to_read>: sar keeps one month of data.
####    Specific file to read can be specified by date. E.g. /var/log/sysstatsa18: 18th day of the month
#### <start_time>: E.g. 09:00:00
#### <end_time>: E.g. 10:30:00
sar -f <file_to_read> -s <start_time> -e <end_time>

# Check system uptime
uptime
```

#### Measuring CPU Activity
```shell
# top: easily get an overview of CPU usage
## To monitor CPU usage of a single process, specify the PID using the -p flag
## -H flag lists other threads that are related to the process. E.g. all apache server processes
top -H -p <pid>

# ApacheBench: Can be used to send simultaneous requests to an an apache server for load testing
ab -n <no_of_requests_per_user> -c <no_of_concurrent_users> <server_url>

# ps: list processes and PIDs
## List processes of the specified user
ps --user <username>
## List all processes 'a', and their users 'u', even if they are not attached to a tty 'x'
ps aux | grep <to_find>

# pstree: show process tree that is tied to a process
## Show a tree of all processes and their relationships
pstree
## Show process tree from a specific process (as root node)
pstree <pid>

# pmap: view the memory map for a process. See all the other process and libraries, it's touching
sudo pmap <pid>

# lsof: check what processes are using a particular file, 
#  or what files a process is using
## What processes are using a particular file
lsof <path_to_file>
## What files are being used by a particular process
lsof -p <pid>

# mpstat: get real-time processor statistics
## ctrl+c at the end gives an average of usage statistics
mpstat <interval_in_seconds>
```

#### Measuring Memory Usage
```shell
# free: get an overview of physical and virtual memory comsumption
## -h: human-readable format
free -h
## run at intervals of 3 seconds
free -s 3

# sar: historical memory usage (monitoring)
## -r: RAM (physical memory); -B: page in/out
sar -r
sar -B
## -S: Swap; -W: Page swap in/out
sar -S
sar -W

# vmstat: get information about physical/virtual memory
## bi: blocks in: disk -> RAM -- always good :-)
## bo: block out: RAM -> -- good or bad: maybe we ran out of RAM? or RAM is being cleaned.
##    bo + si: data moving RAM -> swap -- BAD!
## Get information in megabytes at 3 seconds intervals
vmstat --unit M 3
```

#### Measuring Disk Activity
```shell
# iotop: get overview of IO metrics
## necessary to run as admin (i.e., sudo)
sudo iotop
## monitor accumulative real-time IO: 
##    accumulate IO per process since start of iotop utility
sudo iotop -a
## lsof: list open files
## -p: process id: list open files for specifies pid
lsof -p <pid>

# iostat: show individual disk performance
iostat
## -x: detailed stats
## -t <secs>: refresh every second
## r_await / w_await: shows that read/write operations had to wait; 
##    this is usually a problem and needs to be investigated
iostat -x -t <sec>

# sar: historical disk usage
## rtps: number of file read requests to disk
## wtps: number of file write requests to disk
## bread/s: number of blocks read from disk per second
## bwrtn/s: number of blocks written to disk per second
sar -b
```
