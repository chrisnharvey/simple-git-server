# Simple Git Server

A simple git SSH server, with CLI tool for managing repos.

## Getting Started

```bash
docker run -d -p 22:22 -v /path/to/repos:/git/repos -v /path/to/keys:/git/keys -v /path/to/config:/etc/ssh chrisnharvey/simple-git-server:latest
```

## The Admin CLI

If you SSH into Simple Git Server using the `admin` user, you will be presented with a simple shell to manage git repos.

```bash
ssh admin@localhost
```

### Creating a repo

```bash
create-repo myrepo
```

## Cloning a repo

```bash
git clone git@localhost:myrepo
```
