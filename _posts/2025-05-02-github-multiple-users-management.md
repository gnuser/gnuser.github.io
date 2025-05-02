---
key: github-multiple-users-management
title: github multiple users management
date: 2025-05-02 18:17:43 +0800
typora-root-url: ~/Work/blog/gnuser.github.io
---

Managing Multiple GitHub Accounts Locally

<!--more-->

## Problem

When you need to use multiple GitHub accounts on the same computer, you'll face challenges with SSH keys and authentication.

## Solution: SSH Keys Method

### 1. Generate Different SSH Keys

```bash
# For your personal account
ssh-keygen -t ed25519 -C "personal@example.com" -f ~/.ssh/id_personal

# For your work account
ssh-keygen -t ed25519 -C "work@example.com" -f ~/.ssh/id_work
```

### 2. Add Keys to SSH Agent

```bash
ssh-add ~/.ssh/id_personal
ssh-add ~/.ssh/id_work
```

### 3. Configure SSH

Create or edit `~/.ssh/config`:

```
# Personal GitHub account
Host github.com-personal
   HostName github.com
   User git
   IdentityFile ~/.ssh/id_personal
   IdentitiesOnly yes

# Work GitHub account
Host github.com-work
   HostName github.com
   User git
   IdentityFile ~/.ssh/id_work
   IdentitiesOnly yes
```

### 4. Add SSH Keys to GitHub Accounts

Copy the public keys and add them to the respective GitHub accounts:

```bash
cat ~/.ssh/id_personal.pub
cat ~/.ssh/id_work.pub
```

### 5. Configure Git for Each Repository

```bash
# For personal projects
git config user.name "Personal Name"
git config user.email "personal@example.com"

# For work projects
git config user.name "Work Name"
git config user.email "work@example.com"
```

### 6. Clone Repositories with Correct Host

```bash
# For personal repositories
git clone git@github.com-personal:username/repo.git

# For work repositories
git clone git@github.com-work:company/repo.git
```

## Alternative: HTTPS with Credential Manager

If you prefer HTTPS over SSH:

1. Configure credential manager to store credentials
2. Use different credentials for different repositories
3. Use Git's credential context for different directories

```bash
git config --global credential.useHttpPath true
```

## Testing Your Setup

```bash
ssh -T git@github.com-personal
ssh -T git@github.com-work
```
