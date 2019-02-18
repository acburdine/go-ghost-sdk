workflow "Golang checks" {
  resolves = [
    "sjkaliski/go-github-actions/fmt@v0.2.0",
    "sjkaliski/go-github-actions/lint@v0.2.0",
    "go vet",
    "go test",
  ]
  on = "pull_request"
}

action "sjkaliski/go-github-actions/fmt@v0.2.0" {
  uses = "sjkaliski/go-github-actions/fmt@v0.1.0"
  secrets = ["GITHUB_TOKEN"]
}

action "sjkaliski/go-github-actions/lint@v0.2.0" {
  uses = "sjkaliski/go-github-actions/lint@v0.2.0"
  secrets = ["GITHUB_TOKEN"]
}

action "go vet" {
  uses = "./.github/actions/govet"
  secrets = ["GITHUB_TOKEN"]
}

action "go test" {
  uses = "./.github/actions/gotest"
  secrets = ["GITHUB_TOKEN"]
}
