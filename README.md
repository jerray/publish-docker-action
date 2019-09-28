# Publish Docker Action

[![GitHub Action](https://github.com/jerray/publish-docker-action/workflows/Main/badge.svg)](https://github.com/jerray/publish-docker-action/actions?workflow=Main)
[![codecov](https://codecov.io/gh/jerray/publish-docker-action/branch/master/graph/badge.svg)](https://codecov.io/gh/jerray/publish-docker-action)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/jerray/publish-docker-action?logo=github)](https://github.com/jerray/publish-docker-action/releases)
[![Docker Pulls](https://img.shields.io/docker/pulls/jerray/publish-docker-action?logo=docker)](https://hub.docker.com/r/jerray/publish-docker-action)

Publish Docker Action builds, creates tags and pushes docker image to your docker registry.

## Usage

This simple example uses `Dockerfile` in your workspace to build image, attach the `latest`
tag and push to docker default registry (docker.io). Repository name is your GitHub repository
name by default.

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

Use `file` and `path` arguments to set docker build file or build context if they are not in the default workspace.

### Set up registry and repository name

You can set docker registry with `registry` argument. Change docker repository name with `respository` argument.
For example:

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
```

### Tags

#### Static Tag List

You can use static tag list by providing `tags` argument. Concat multiple tag names with commas.

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
    tags: latest,newest,master
```

This example builds the image, creates three tags, and pushes all of them to the registry.

* `jerray/publish-docker-action:latest`
* `jerray/publish-docker-action:newest`
* `jerray/publish-docker-action:master`

#### Auto Tag

This action can generate image tags automatically by different `refs` types:

* branch: uses the branch name as docker tag name (`master` branch is renamed to `latest`).
* pull request: attaches a `pr-` prefix to branch name asdocker image tag. To allow pull request build, you must set `with.allow_pull_request` to `true`.
* tag: checks if the tag name is valid semantic version format. If not, it uses git tag name as docker image tag directly. Else it generates three tags based on the version number,each followed with pre-release information.

Examples:

| Git | Docker Tag |
| --- | --- |
| branch `master` | `latest` |
| branch  `2019/09/28-new-feature` | `2019-09-28-new-feature` (`/` is replaced to `-`) |
| pull request `master` | `pr-master` |
| tag `1.0.0` | `1`, `1.0`, `1.0.0` |
| tag `v1.0.0` | `1`, `1.0`, `1.0.0`  (prefix `v` is allowed) |
| tag `v1.0.0-rc1` | `1-rc1`, `1.0-rc1`, `1.0.0-rc1` |
| tag `20190921-actions` | `20190921-actions` (not semantic version) |

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
    auto_tag: true
```

Auto tagging will override `with.tags` list.

### Cache

Provide `with.cache` argument to build from cache.

### Build Args

Use `with.build_args` to provide docker build-time variables. Multiple variables must be separated by comma. 

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
    build_args: HTTP_PROXY=http://127.0.0.1,USER=nginx
```

## Note

Please use the latest released version rather than master.
