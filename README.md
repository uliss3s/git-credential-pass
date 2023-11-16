# git-credential-pass
Go implementation of a git custom credential helper using the standard unix password manager (pass) store

Obs: This helper implements only the "get" operation.

## Getting started

---
### pass
Configure the Password Store in the following structure: hostname/username

Ex.: pass insert github.com/me@email.com

### git
Configure .gitconfig. Ex.:
```
[credential "http://mylocalgitlabinstance.local"]
    username = me@email.com
    helper = pass

[credential "https://github.com"]
    username = me@email.com
    helper = pass
```

## Docs

---
Git Credentials: https://git-scm.com/docs/gitcredentials

Unix standard password manager: https://www.passwordstore.org/