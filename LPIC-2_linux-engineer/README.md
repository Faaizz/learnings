# LPIC-2 Linux Engineer Cert Prep

Notes from LPIC-2 Linux Engineer (201-450) Cert Prep video course on LinkedIn Learning.
*Note: A debian-based linux distro is assumed for the listed commands.*

## Content
- [LPIC-2 Linux Engineer Cert Prep](#lpic-2-linux-engineer-cert-prep)
  - [Content](#content)
  - [Supporting Linux](#supporting-linux)
    - [Capacity Planning](#capacity-planning)

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
