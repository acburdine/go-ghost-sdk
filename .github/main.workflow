workflow "Golang checks" {
  on = "push"
  resolves = ["sjkaliski/go-github-actions/fmt@v0.1.0"]
}

action "sjkaliski/go-github-actions/fmt@v0.1.0" {
  uses = "sjkaliski/go-github-actions/fmt@v0.1.0"
  secrets = ["GITHUB_TOKEN"]
}
