#!/bin/bash
#
#    This script creates a new blog post with metadata in ./_posts
#    folder. Date will be generated according to the current time in
#    the system. Usage:
#
#        ./newpost.sh "My Blog Post Title"
#

typorarooturl="/Users/chenjing/workspace/github/gnuser.github.io"

title=$1

if [[ -z "$title" ]]; then
    echo "usage: newpost.sh \"My New Blog\""
    exit 1
fi

bloghome=$(cd "$(dirname "$0")"; pwd)
url=$(echo "$title" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')
filename="$(date +"%Y-%m-%d")-$url.md"
filepath=$bloghome/_posts/$filename

if [[ -f $filepath ]]; then
    echo "$filepath already exists."
    exit 1
fi

touch $filepath

echo "---" >> $filepath
echo "key: ${url}" >> $filepath
echo "title: ${title}" >> $filepath
echo "date: $(date +"%Y-%m-%d %H:%M:%S %z")" >> $filepath
echo "typora-root-url: ${typorarooturl}" >> $filepath
echo "---" >> $filepath
echo "" >> $filepath
echo "<!--more-->" >> $filepath


echo "Blog created: $filepath"