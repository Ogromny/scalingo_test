# Canvas for Backend Technical Test at Scalingo

## Instructions

* From this canvas, respond to the project which has been communicated to you by our team
* Feel free to change everything

## Execution

```
docker-compose up
```

or
```
podman-compose up
```

Application will be then running on port `5000`

## API

### repos

> Allow to search on github's repositories

url: `/repos`
params:
    - stars: unsigned (support basic condition, i.e. `>400`, `<200`)
    - language: string
    - user: string
    - org: string
    - size: unsigned
    - forks: unsigned
    - license: string
    - archived: bool

#### Example
```
$ curl localhost:5000/repos&stars=>500&stars=<500&language=c
{ ... }
```

### stats

> Show a short resume about every previous search

url: `/stats`
params: none

#### Example
```
$ curl localhost:5000/stats
{ ... }
```
