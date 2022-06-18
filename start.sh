#!/bin/sh

/usr/bin/ssh-keygen -A
/usr/sbin/sshd -f /etc/sshd_config -D