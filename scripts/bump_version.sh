#!/bin/bash
set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 [patch|minor|major]"
  exit 1
fi

CURRENT_VERSION=$(git describe --tags --abbrev=0 | sed 's/v//')
IFS='.' read -r -a parts <<< "$CURRENT_VERSION"

major=${parts[0]}
minor=${parts[1]}
patch=${parts[2]}

case "$1" in
  patch)
    patch=$((patch + 1))
    ;;
  minor)
    minor=$((minor + 1))
    patch=0
    ;;
  major)
    major=$((major + 1))
    minor=0
    patch=0
    ;;
  *)
    echo "Invalid bump type. Use patch, minor, or major."
    exit 1
    ;;
esac

NEW_VERSION="v${major}.${minor}.${patch}"

echo "Bumping version: $CURRENT_VERSION -> $NEW_VERSION"

git tag $NEW_VERSION
git push origin $NEW_VERSION
