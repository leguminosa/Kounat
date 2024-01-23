#!/bin/bash

# Detect the current branch.
current_branch=$(git rev-parse --abbrev-ref HEAD | awk '{print tolower($0)}')

# Check if the current branch is main.
if [[ $current_branch == "main" ]]; then
    printf "\e[0mGenerating mockgen on \e[36mmain\e[0m branch is not allowed. Please checkout to a new branch and try again.\e[0m\n"
    exit 1
fi

# Get the parent commit of the current branch.
parent_commit=$(git log --grep="^Merge pull request #" --max-count=1 --pretty=format:"%H")

# Get the list of files that are changed in the current branch.
# It will also include all untracked files, as long as they have already been staged.
# The first git diff command will capture all changes that have been commited
# to the current branch since diverging from the parent commit.
# The second git diff command will include all changes (including
# all untracked files) that have been staged, but not commited yet.
# And finally, the third git diff will include all modified files
# in working directory that are still not yet staged.
temp_file=git_diff.out
touch "$temp_file"
git diff "$current_branch".."$parent_commit" --name-only >> "$temp_file"
git diff --name-only --cached >> "$temp_file"
git diff --name-only >> "$temp_file"
files=$(sort "$temp_file" | uniq)
rm "$temp_file"

# Iterate through all affected files.
for file in $files; do
    # Make sure file exists and is not a directory.
    if [[ ! -f $file ]]; then
        continue
    fi

    # Only process golang files.
    if [[ ! $file == *.go ]]; then
        continue
    fi

    # Skip golang test files.
    if [[ $file == *_test.go ]]; then
        continue
    fi

    # Skip golang mock files.
    if [[ $file == *.mock.go ]]; then
        # Remove mock file if there is no actual file.
        if [[ ! -f ${file%.mock.go}.go ]]; then
            git rm -q "$file"
            printf "\e[0m%-50s\e[35m%s\e[0m\n" "$file" "removed"
        fi
        continue
    fi

    # Find for exported interface inside the file
    # by using a simple interface keyword grep.
    if [[ ! $(grep -E "type [A-Z][A-Za-z0-9_]* interface" "$file" | wc -l) -gt 0 ]]; then
        continue
    fi

    # Manipulate the file name to get the mock file name.
    # Can either use sed command or bash string manipulation.
    # mock_file=$(sed 's/\.go$/\.mock\.go/' <<< "$file")
    mock_file=${file%.go}.mock.go
    dir=$(dirname "$file")

    # Generate mock file using mockgen command.
    $(go env GOPATH)/bin/mockgen -source=${file} -destination=${mock_file} -package=${dir##*/}
    git add $mock_file
    printf "\e[0m%-50s\e[34m%s\e[0m\n" "$mock_file" "generated"
done
