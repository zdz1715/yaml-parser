# yaml-parser
处理`yaml`、`yml`文件的自定义注解解析

# 使用

```shell
$ yaml-parser --help
Handles custom annotation parsing of YAML files

Usage:
  yaml-parser [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  param       Extract parameters from a single yaml file
  split       Dividing a YAML file into multiple

Flags:
  -h, --help   help for yaml-parser

Use "yaml-parser [command] --help" for more information about a command.

```

## 支持的注解
- `# +parser:tag=value`

  适用于`split`
- `# +parser:param:key1=value1,key2=value2`

  适用于`param`
## 功能
`deploy-helm.yaml`（例子）
```yaml
# +parser:tag=stage
# +parser:tag=prod
# +parser:param:namespace=default,release_name=app-1
name: deploy-helm-stage
tag:
  - 111
  - 222
---
# +parser:param:namespace=default-1,release_name=app-2
name: deploy-helm-1
tag:
  - 333
  - 444
```
### 分割`split`
> 依赖于yaml文件的分隔符`---`
```shell
$ yaml-parser split deploy-helm.yaml
[deploy-helm.yaml] deploy-helm_stage.yaml deploy-helm_prod.yaml deploy-helm_1.yaml
```
- `deploy-helm_stage.yaml`
```yaml
# +parser:tag=stage
# +parser:tag=prod
# +parser:param:namespace=default,release_name=app-1
name: deploy-helm-stage
tag:
  - 111
  - 222
```
- `deploy-helm_prod.yaml`
```yaml
# +parser:tag=stage
# +parser:tag=prod
# +parser:param:namespace=default,release_name=app-1
name: deploy-helm-stage
tag:
  - 111
  - 222
```
- `deploy-helm_1.yaml`
```yaml
# +parser:param:namespace=default-1,release_name=app-2
name: deploy-helm-1
tag:
  - 333
  - 444
```

### 自定义参数 `param`
```shell
# 获取所有参数
$ yaml-parser param deploy-helm.yaml
[{"key":"namespace","value":"default"},{"key":"release_name","value":"app-1"},{"key":"namespace","value":"default-1"},{"key":"release_name","value":"app-2"}]
# 获取特定key的参数，相同key后面值会覆盖前面
$ yaml-parser param deploy-helm.yaml --key release_name
app-2
```
