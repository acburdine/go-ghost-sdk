FROM golang:1.11

LABEL "com.github.actions.name"="go vet"
LABEL "com.github.actions.description"="Run go vet"
LABEL "com.github.actions.icon"="terminal"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="https://github.com/acburdine/go-ghost-sdk"
LABEL "homepage"="https://github.com/acburdine/go-ghost-sdk"
LABEL "maintainer"="Austin Burdine <acburdine@gmail.com>"

RUN apt-get update && \
    apt-get install -y jq

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
