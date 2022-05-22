# Foundations of Ansible

## Content
- [Foundations of Ansible](#foundations-of-ansible)
  - [Content](#content)
  - [Introduction to IT Automation](#introduction-to-it-automation)
    - [IT Automation Tools](#it-automation-tools)
    - [Ansible Concepts](#ansible-concepts)
    - [Ansible Inventory](#ansible-inventory)
    - [Ansible Playbooks](#ansible-playbooks)
    - [Ansible Tower](#ansible-tower)
    - [Ansible Provisioning](#ansible-provisioning)
    - [Ansible Configuration Management](#ansible-configuration-management)
  - [Getting Started with Ansible](#getting-started-with-ansible)
    - [Control Node Best Practices](#control-node-best-practices)
    - [Managed Nodes](#managed-nodes)
    - [Configuring Ansible](#configuring-ansible)
    - [Useful Commands](#useful-commands)
    - [Ansible Ad-hoc Commands](#ansible-ad-hoc-commands)
      - [Syntax](#syntax)
      - [Patterns](#patterns)
      - [Managing Files and Packages](#managing-files-and-packages)
      - [Managing Services](#managing-services)
      - [Managing Users](#managing-users)
      - [Gathering Data](#gathering-data)
    - [Ansble Command Line Tools](#ansble-command-line-tools)
      - [Manage Connection Methods](#manage-connection-methods)
      - [The `ansible` Command](#the-ansible-command)
      - [View Ansible Configuration](#view-ansible-configuration)
      - [Show Ansible Inventory](#show-ansible-inventory)
    - [Help Pages](#help-pages)
    - [Fedora Tips](#fedora-tips)
    - [Virtualbox Tips](#virtualbox-tips)
  - [Ansible Playbooks](#ansible-playbooks-1)
  - [References](#references)

## Introduction to IT Automation
### IT Automation Tools
- Chef
- Puppet
- Salt
- Ansible Engine

### Ansible Concepts
- Managed over SSH Connections
- Supports SSH pipelining
- Inventory describes location of nodes
- Stores sensitive information in Ansible Vault
- Nodes are agentless
- Any host can be the control node, does not require a central server
- Installs and runs temporary modules on nodes
- Modules communicate via STDIN/STDOUT with a JSON based protocol
- Ansible Playbooks strive to be idempotent

### Ansible Inventory
- Stored as INI or YAML files
- Hosts can be specified using IP addresses or hostnames
- Hosts can be specified in groups, or group of groups
- Can be managed dynamically using plugins or scripts
- Host groups can be set up along: what, where, and when. E.g.:
    ```ini
    # What
    [webservers]
    server1.local.net
    server2.local.net
    [dbservers]
    server3.local.net
    # Where
    [seattle]    
    server1.local.net
    server3.local.net
    [frankfurt]
    server2.local.net
    # When
    [prod]
    server1.local.net
    [dev]
    server2.local.net
    server3.local.net
    ```
- Help on setting up inventories can be found at: [How to build your inventory- Ansible Documentation](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html)

### Ansible Playbooks
Configurations are stored in Playbooks. A Playbook is a YAML file that expresses:
- Configuration
- Deployments
- Orchestration

They allows Ansible to perform operation on Nodes.
A Playbook maps groups of hosts to sets of roles. 
A role is represented by a set of Ansible tasks.

### Ansible Tower
- REST API
- Web service
- Web console
- Hub to manage Ansible installations
- Commercial RedHat product
- Open source version: AWX

### Ansible Provisioning
Ansible is capable of provisioning:
- Bare-metal servers
- Virtualized systems
- Cloud systems
- Storage systems
- Network devices

### Ansible Configuration Management
- Relies on OpenSSH and only requires an SSH key or password to manage a system
- Configuration describes a desired state
- Allows for a "dry run" execution
- Delivers and executes modules remotely
- Does not require root privileges but supports sudo


## Getting Started with Ansible
### Control Node Best Practices
- Should be near managed nodes. If nodes are local, control node should be local; if nodes are in the cloud, control node should be in the cloud.
- Different pugins may have different requirements
- Number of file handles may need to be increased in MacOS
### Managed Nodes
Requires:
- \>= python2.4
- if == python2.4, then `python-simplejson` is required
- looks for python at `/usr/bin/python`
- SELinux systems require `libselinux-python` package
- Ansible raw/script moduless do not require python, thus can be used to install python on managed nodes

### Configuring Ansible
Ansible can be configured via various sources, including (listed from lowest to highest priority):
- `.ansible,cfg`
- Environment variables
- Command line options
- Playbook keywords
- Ansible variables: Set in playbooks using `vars:` or in the CLI using `-e` flag
 
`.ansible.cfg` file is searched in the following order, listed from highest to lowest priority:
- Location in `$ANSIBLE_CONFIG` environment variable
- `ansible.cfg` in the current directory
- `~/.ansible.cfg`: User home directory (yes `.ansible.cfg`- hidden file, not a typo)
- `/etc/ansible/ansible.cfg`

### Useful Commands
```shell
# Check ansible version
ansible --version

# Set inventory file using the -i flag
# List all available hosts
ansible -i /path/to/hosts all --list-hosts 

# -m flag can be used to send ansible ad-hoc commands to hosts
# Check communication with hosts
ansible -i /path/to/hosts all -m ping

# -a flag can be used to send regular linux command to hosts
# Check system uptime
ansible -i /path/to/hosts all -a "uptime"

# List avaibale ansible plugins
ansible-doc -l
# -t flag can be used to specify plugin type
# Plugin types include: become, cache, callback, cliconf, connection, 
#   httpapi, inventory, lookup, netconf, shell, strategy, vars
ansible-doc -l -t cache
# Info about specific plugin
ansible-doc file
# Snippet about specific plugin. 
# Returns just the available options and their description
ansible-doc -s file
```

### Ansible Ad-hoc Commands
Ansible ad-hoc commands can use any of the ansible modules and are useful for:
- Automating single tasks
- testing commands before including them in playbooks

#### Syntax
```shell
# *Note: the arrangement of the arguments is not strict*
ansible [ansible flags] [pattern] -m [module] -a "[module options]"
# To run commands directly in the hosts configured shell, use the shell module
# For command string that have a variable in them, single quotes must be used.
ansible all -m shell -a 'echo $PWD'
# To run simpler commands, we can use the command module by omitting the -m flag
ansible all -a "/sbin/reboot"
# Ansible by default spins up 5 processes at a time, -f flag can be used to set this per command
ansible all -a "/sbin/reboot" -f 10
# User to run a command as on the host can be specified with the -f flag
ansible all -a 'echo $PWD' -u user1
# -b and -K flags allow prvilege escalation to root and specification of user's password respectively
ansible all -a "/sbin/reboot" -b -K
```
The shell module:
- Runs commands in shell
- is more complex but more compatible
- support pipes and redirects
- support environment variables

The command module:
- Doesn't use a shell
- is simple and secure
- doesn't support pipes and redirects
- doesn't support environment variables

#### Patterns
Patterns are used to include/exclude subsets of hosts, and can use wildcards or regular expressions. They can match:
- A single host
- an IP address
- an inventory group
- a set of groups
- all hosts

Some common patterns include:
- `all` or `*` for all hosts
- `host-name` for a specific host
- `group-name` for a group of hosts
- `g1:g2` for all hosts that match `g1` or `g2`
- `g1:!g2` for all hosts that match `g1` excluding those that match `g2`
- `g1:&g2` for all hosts that are in both `g1` and `g2`
Patterns can be combined for more complex matches such as: `webservers:dbservers:&test:!prod`

#### Managing Files and Packages
```shell
# Touch a file
ansible -i $PATH_TO_HOSTS -m file -a 'path=~/file1.txt state=touch mode=700' rhhost1
# Change file permissions
ansible -i $PATH_TO_HOSTS -m file -a 'path=~/file1.txt mode=600' rhhost1
# Change file owner/group
ansible -i $PATH_TO_HOSTS -m file -a 'path=/home.user1/file1.txt mode=600 owner=root group=root' -b -K rhhost1
# Create a directory
ansible -i $PATH_TO_HOSTS -m file -a 'dest=/home/user1/newdir state=directory mode=700' rhhost1
# Delete directory recursively
ansible -i $PATH_TO_HOSTS -m file -a 'dest=/home/user1/newdir state=absent mode=700' rhhost1

# Create a text file with some content
# force=no means the contents of the file wouldn't be overwritten if file already exists
ansible -i $PATH_TO_HOSTS -m copy -a "dest=~/file2.txt content='file content' force=no mode=700" rhhost1

# Install package with yum module
# state=installed and state=present install the package
# state=latest installs/updates the package to latest version
# state=removed or state=absent uninstalls the package
ansible -i $PATH_TO_HOSTS -m yum -a "name=httpd state=installed" -b -K rhhost1
# Update package list
ansible -i $PATH_TO_HOSTS -m yum -a "name=* state=latest" -b -K rhhost1
# List installed packages
ansible -i $PATH_TO_HOSTS -m yum -a "list=installed" rhhost1
# yum group install
ansible -i $PATH_TO_HOSTS -m yum -a "name='@Development tools' state=installed" -b -K rhhost1
```

#### Managing Services
```shell
# Use the services module to start httpd
# state=reloaded reloads the service
# state=stopped stops the service
# enabled=yes to enable the service. i.e. start on boot
# disabled=yes to disable the service. i.e. do not start on boot
ansible -i $PATH_TO_HOSTS -m service -a "name=httpd state=started" -b -K rhhost1
# Check service status. 
# Returns *"changed": false* if current state matches state that's being checked
# check started
ansible -i $PATH_TO_HOSTS -m service -a "name=httpd state=started" rhhost1 --check
# check enabled
ansible -i $PATH_TO_HOSTS -m service -a "name=httpd enabled=true" rhhost1 --check

# systemctl can also be run on the host via shell/command to manage services
ansible -i $PATH_TO_HOSTS -m shell -a "systemctl status httpd" rhhost1
```

#### Managing Users
```shell
# Add a new user using the user module
ansible -i $PATH_TO_HOSTS -m user -a "name=user2 home=/home/user2 shell=/bin/bash state=present" -b -K rhhost1

# Change user's primary group
ansible -i $PATH_TO_HOSTS -m user -a "name=user2 group=user2" -b -K rhhost1
# Add a supplementary group to user
ansible -i $PATH_TO_HOSTS -m user -a "name=user2 groups=wheel" -b -K rhhost1

# Delete user
ansible -i $PATH_TO_HOSTS -m user -a "name=user2 state=absent" -b -K rhhost1

# Add user group
ansible -i $PATH_TO_HOSTS -m group -a "name=group2 state=present" -b -K rhhost1
```

#### Gathering Data
```shell
# Get all facts about a host
ansible -i $PATH_TO_HOSTS -m setup rhhost1
# To extract a subset of info
# Subsets include: all, network, min, hardware, virtual, facter
ansible -i $PATH_TO_HOSTS -m setup -a 'gather_subset=network' rhhost1
# ! can be used to negate subsets
# To get only hardware facts, run:
ansible -i $PATH_TO_HOSTS -m setup -a 'gather_subset=!all,!min,hardware' rhhost1
# To even fine-grain the info further and return ospecified section, we can filter
ansible -i $PATH_TO_HOSTS -m setup -a 'gather_subset=!all,!min,hardware filter=ansible_device_links' rhhost1
# Dump into file
ansible -i $PATH_TO_HOSTS -m setup rhhost1 --tree /path/to/file
```

An index of all ansible modules can be found at [Module Index- Ansible Documentation](https://docs.ansible.com/ansible/2.9/modules/modules_by_category.html).

### Ansble Command Line Tools
#### Manage Connection Methods
Ansibble by default uses OpenSSh for SSH connections, if this is not available, it falls back to `paramiko` which is a Python implementation of OpenSSH.

To authenticate SSH without keys (i.e. with a password) use `--ask-password` or `-k` flag on the CLI.

To specify a private key, use the `ansible_ssh_private_key_file` option in the Ansible inventory or:
```shell
ansible -i $PATH_TO_HOSTS -m <module> -a "<options>" --private-key /path/to/key <hosts>
```
on the CLI.

Host key checking adds an extra layer of security by maintaining an ID database of hosts which it checks against to prevent man-in-the-middle attacks. However, this might result in increased maintenance (`known_hosts` file might need to be modified in some cases). To disable host key checking, set `host_key_checking=False` in `ansible.cfg` or `export ANSIBLE_HOST_KEY_CHECKING=False` on the control host.

Excluding SSH, there are other connections methods available to Ansible, these include:
- local: lets you run Playbooks on the control host itself
- docker: lets you run Playbooks in docker containers
- kubectl: lets you run Playbooks in kubernetes pods
- vmware_tools: lets you run Playbook on VMWare hosts
- podman, etc.

To get a full list, run:
```shell
ansible-doc -t connection -l
# Get more info about a single cpnnection method with
ansible-doc -t connection <plugin-name>
```

#### The `ansible` Command
Is a simple and short way to get things done. Some available options (flags) include:
- `--list-host`: lists matching hosts and do nothing else
- `--private-key`: specify private key
- `--syntax-check`: check playbook syntax
- `--playbook-dir`: specify playbook directory 
- `-e`: set environment variables
- `-i`: specify inventory file
- `-m`: specify Ansible module to use
- `-a`: provide arguments for Ansible model in use
- `-v`: set verbose mode

#### View Ansible Configuration
Ansible configuration can be viewed using the `ansible-config` command.
```shell
# List available config
ansible-config list
# To view current config
ansible-config dump
# Use the --only-changed flag to view configs that defer from the default
ansible-config dump --only-changed
```

#### Show Ansible Inventory
```shell
# List hosts categorized by groups
# use the -y option to output YAML instead of JSON
ansible-inventory --list
# output to file with --output flag
ansible-inventory --list --output /path/to/file
# get ASCII graph with --graph option
ansible-inventory --graph
```

### Help Pages
There are man pages available for:
- `ansible-config`
- `ansible-console`
- `ansible-galaxy`
- `ansible-inventory`
- `ansible-playbook`
- `ansible-pull`
- `ansible-vault`
- `ansible-doc`

### Fedora Tips
```shell
# Install Development tools
sudo yum group install -y "Development tools"

# Change hostname
sudo hostnamectl set-hostname <new-hostname>

# Install argcomplete for tab completion
sudo yum install -y python3-argcomplete
# Activate argcompelte
sudo activate-global-python-argcomplete3
# or on some systems
sudo activate-global-python-argcomplete

# Create ssh key
ssh-keygen
# Copy ssh key over to a different host
ssh-copy-id
```

### Virtualbox Tips
- Setup a `host-only-adapter` network to allow the VMs communicate with one another and the host.
  File > Host Network Manager

- Setup a `NAT` network to allow the VMs access the internet. 


## Ansible Playbooks
Below is a basic playbook for managing an apache http server:
```yaml
# Playbook name
name: Apache server installed
# Hosts to match
hosts: webservers
# Get root privileges
become: yes
# Tasks (Plays)
tasks:
  - name: Latest version installed
    yum:
      name: httpd
      state: latests
  - name: Server enabled and running
    service:
      name: httpd
      state: started
      enabled: yes
```
After setting up a playbook, it is often useful to check for syntax errors, this can be done with:
```shell
ansible-playbook --syntax-check /path/to/playbook
```
And run playbook with:
```shell
ansible-playbook /path/to/playbook -b -K
```

## References
- [Red Hat Certified Engineer (EX294) Cert Prep: 1 Foundations of Ansible- LinkedIn Learning](https://www.linkedin.com/learning/red-hat-certified-engineer-ex294-cert-prep-1-foundations-of-ansible)
- [Red Hat Certified Engineer (EX294) Cert Prep: 2 Using Ansible Playbooks- LinkedIn Learning](https://www.linkedin.com/learning/red-hat-certified-engineer-ex294-cert-prep-2-using-ansible-playbooks)
