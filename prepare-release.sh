#!/bin/sh

# Execute from repository root with a clean working tree!

if test $# -ne 1
then
	echo "usage: ./prepare-release.sh x.x.x"
    exit 1
fi
PRDFLG='config/productionFlags.go'
git checkout -b "rc-$1" develop
sed "s/var Production =.*/var Production = true \/\/ v$1 (production)/" "$PRDFLG" > "$PRDFLG.tmp" && mv "$PRDFLG.tmp" "$PRDFLG"
git add "$PRDFLG"
git commit -m 'Set flag Production = true'

git checkout develop
sed "s/var Production =.*/var Production = false \/\/ v$1 (development)/" "$PRDFLG" > "$PRDFLG.tmp" && mv "$PRDFLG.tmp" "$PRDFLG"
git add "$PRDFLG"
git commit -m "Diverge $PRDFLG"
