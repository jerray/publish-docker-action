# Publish Docker Action

Publish Docker Action is used to build, tag and publish docker image to your docker registry.

## Usage

This simple example will use `Dockerfile` in your workspace to build image, attach the `latest`
tag and push to docker default registry (docker.io). Repository name is your GitHub repository
name by default.

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

Use `file` and `path` arguments to set docker build file or build context if they are not placed
in the default workspace direcotry.

### Set up registry and repository name

Registry and repository name can be changed with `registry` and `repository` arguments. For example:

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

You can use static tag list by specify `tags` arguments. Tag names must be separated by comma.

```yaml
- uses: jerray/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: jerray/publish-docker-action
    tags: latest,newest,master
```

Example above will build image and create three tags, and push all of them to the registry.

* `jerray/publish-docker-action:latest`
* `jerray/publish-docker-action:newest`
* `jerray/publish-docker-action:master`

#### Auto Tag

This action can generate image tag automatically base on the different `refs` type.

If the `refs` refers to a branch, it uses the branch name as docker image name (`master` branch is renamed to `latest`).

If the `refs` refers to a pull request, it attaches a `pr-` prefix to branch name as
docker image tag. To allow pull request build, you must set `with.allow_pull_request` to `true`.

When `refs` refers to a tag, it checks if the tag name is valid semantic version. If not, it uses
tag name as docker image tag directly. Else it generates three tags based on the version number,
each followed with pre-release information if there is any. For example:

* git tag `1.0.0` is mapped to `1`, `1.0`, `1.0.0`
* git tag `v1.0.0` is the same as above (prefix `v` is allowed)
* git tag `v1.0.0-rc1` is mapped to `1-rc1`, `1.0-rc1`, `1.0.0-rc1`
* git tag `20190921-actions` is not a valid semantic version string, so docker image tag is just the same

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

Cache image can be used to build image. Just provide `with.cache` argument.

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
