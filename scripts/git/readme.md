# git related scripts

- `commit-msg` enforces commit style
- `commit-msg.template` is template for enforces commit style

```shell
git config commit.template $(pwd)/scripts/git/commit-msg.template
cp ./scripts/git/commit-msg .git/hooks/commit-msg
```
