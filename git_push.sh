#!/bin/sh
# ref: https://help.github.com/articles/adding-an-existing-project-to-github-using-the-command-line/
#
# Usage example: /bin/sh ./git_push.sh wing328 openapi-petstore-perl "minor update" "gitlab.com"

git_user_id=$1
git_repo_id=$2
git_token=$3
git_host=$4


if [ "$git_host" = "" ]; then
    git_host="github.com"
    echo "[INFO] No command line input provided. Set \$git_host to $git_host"
fi

if [ "$git_user_id" = "" ]; then
    git_user_id="GIT_USER_ID"
    echo "[INFO] No command line input provided. Set \$git_user_id to $git_user_id"
fi

if [ "$git_repo_id" = "" ]; then
    git_repo_id="GIT_REPO_ID"
    echo "[INFO] No command line input provided. Set \$git_repo_id to $git_repo_id"
fi

# Initialize the local directory as a Git repository
git init

git add .

# Commits the tracked changes and prepares them to be pushed to a remote repository.
git commit -m "feat: api updated by cicd"

git checkout -b main
git remote add origin https://chenyunda218:"$git_token"@github.com/$git_user_id/$git_repo_id.git
git push -u --force origin main 2>&1 | grep -v 'To https'

