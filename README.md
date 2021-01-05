# TODO to Issue
`ti` is a CLI to convert TODO's in your projects into GitHub issues.

## Installation
### Install with `go`
```
go get https://github.com/rkabani19/ti
```

### Install from source
Clone repository
```
git@github.com:rkabani19/ti.git && cd ti
```
Install `ti`
```
go install
```

## Usage
`Usage: ti [flags] GITHUB_TOKEN`

1. Generate a [personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)
2. Navigate into GitHub project directory
3. Run `ti` with your personal access token, and any flags you desire
4. You will be met with an interactive interface, that will allow you to `Open Issue`, `Skip Issue`, or `Exit` for each TODO found

## Demo
![example](https://user-images.githubusercontent.com/25307996/103675584-1dbe8280-4f4e-11eb-96ce-b4ff8c14a8e6.png)
