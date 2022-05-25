# LPIC-2 Linux Engineer Cert Prep

Notes from LPIC-2 Linux Engineer (201-450) Cert Prep video course on LinkedIn Learning.
*Note: A debian-based linux distro is assumed for the listed commands.*

## Content
- [LPIC-2 Linux Engineer Cert Prep](#lpic-2-linux-engineer-cert-prep)
  - [Content](#content)
  - [Supporting Linux](#supporting-linux)
    - [Capacity Planning](#capacity-planning)
      - [Measuring CPU Activity](#measuring-cpu-activity)

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
