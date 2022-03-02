# yaml-parser
处理`yaml`、`yml`文件的切割、合并等

# 使用

```shell
$ yaml-parser --help
Deal with the cutting and merging of YAML files

Usage:
  yaml-parser [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  split       Dividing a YAML file into multiple

Flags:
  -h, --help   help for yaml-parser

Use "yaml-parser [command] --help" for more information about a command.
```

## 分割
> 兼容yaml文件的分隔符`---`
### 设置tag
1. 创建`big-1.yaml`文件（例子）
```yaml
name: big-1_1
tag:
  - 111
  - 222
---
# +parser:tag=dev
name: big-1_2
tag:
  - 333
  - 444
```
2. 使用
```shell
$ yaml-parser split big-1.yaml
[big-1.yaml] big-1_0.yaml big-1_dev.yaml
```
3. 结果
- `big-1_0.yaml`
```yaml
name: big-1_1
tag:
  - 111
  - 222
```
- `big-1_dev.yaml`
```yaml
# +parser:tag=dev
name: big-1_2
tag:
  - 333
  - 444
```
### 无tag
1. 创建`big-2.yaml`文件（例子）
```yaml
name: big-2_1
tag:
  - one
  - two
---
name: big-2_2
tag:
  - three
  - four
```
2. 使用
```shell
$ yaml-parser split big-2.yaml
[big-1.yaml] big-2_0.yaml big-2_1.yaml
```
3. 结果
- `big-2_0.yaml`
```yaml
name: big-2_1
tag:
  - one
  - two
```
- `big-2_1.yaml`
```yaml
name: big-2_2
tag:
  - three
  - four
```