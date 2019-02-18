#!/bin/sh
set -e

# enable go modules
export GO111MODULE=on

cd "${GO_WORKING_DIR:-.}"

set +e
OUTPUT="$(go test -v -mod=vendor ${GO_TEST_PATHS_:-./...})"
SUCCESS=$?
set -e

# exit if `go vet` passes
if [ $SUCCESS -eq 0 ]; then
  echo "$OUTPUT"
  exit 0
fi

# Post results back as comment.
COMMENT="#### \`go test\`
\`\`\`
$OUTPUT
\`\`\`
"
PAYLOAD=$(echo '{}' | jq --arg body "$COMMENT" '.body = $body')
COMMENTS_URL=$(cat /github/workflow/event.json | jq -r .pull_request.comments_url)
curl -s -S -H "Authorization: token $GITHUB_TOKEN" --header "Content-Type: application/json" --data "$PAYLOAD" "$COMMENTS_URL" > /dev/null

exit $SUCCESS
