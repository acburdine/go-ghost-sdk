workflow "Golang checks" {
  resolves = ["sjkaliski/go-github-actions/fmt@v0.2.0", "sjkaliski/go-github-actions/lint@v0.2.0"]
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
