# Tissue
`tissue` is a CLI to convert TODO's in your projects into GitHub issues.

## Installation
### Install with `go`
```
go get https://github.com/rkabani19/tissue
```

### Install from source
Clone repository
```
git@github.com:rkabani19/tissue.git && cd tissue
```
Install `tissue`
```
go install
```

## Usage
```
tissue [flags] GITHUB_TOKEN
```
1. Generate a [personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)
2. Navigate into GitHub project directory
3. Run `tissue` with your personal access token, and any flags you desire
4. You will be met with an interactive interface, that will allow you to `Open Issue`, `Skip Issue`, or `Exit` for each TODO found

## Demo
![example](https://user-images.githubusercontent.com/25307996/103799233-d0f3ae00-5018-11eb-8765-0f62e5e597bd.png)
