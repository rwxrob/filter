# ð³Go Filter Library

![WIP](https://img.shields.io/badge/status-wip-red.svg)
[![GoDoc](https://godoc.org/github.com/rwxrob/filter?status.svg)](https://godoc.org/github.com/rwxrob/filter)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/filter)](https://goreportcard.com/report/github.com/rwxrob/filter)
[![Coverage](https://gocover.io/_badge/github.com/rwxrob/filter)](https://gocover.io/github.com/rwxrob/filter)

## Usage

```
z filter yt link
filter yt link
z yt link
```

ðº <https://youtu.be/8st1NhaKDCA>

## Example Bash Filters to Convert

This one adds Go badges:

```bash
#!/usr/bin/env bash

url="$1"

if [[ -e go.mod ]];then
  url=$(head -1 go.mod)
  url=${url#* }
fi

echo "![WIP](https://img.shields.io/badge/status-wip-red.svg)"
echo "[![GoDoc](https://godoc.org/$url?status.svg)](https://godoc.org/$url)"
echo "[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)"
echo "[![Go Report Card](https://goreportcard.com/badge/$url)](https://goreportcard.com/report/$url)"
echo "[![Coverage](https://gocover.io/_badge/$url)](https://gocover.io/$url)"
echo
```

Universal way to comments:

```bash
#!/usr/bin/env bash

while IFS= read -r line; do
  echo "${1:-#} $line"
done
```

Detect and replace emojis on save:

```bash
#!/usr/bin/env bash

declare -A emoji
emoji[smile]=ð
emoji[bear]=ð»
emoji[hamburger]=ð
emoji[lightbulb]=ð¡
emoji[idea]=ð¡
emoji[comment]=ð¬
emoji[chat]=ð¬
emoji[pomo]=ð
emoji[stop]=ð
emoji[warning]=â ï¸
emoji[rant]=ð¤¬
emoji[tv]=ðº
emoji[update]=ð°
emoji[tux]=ð§
emoji[tux]=ð§
emoji[facepalm]=ð¤¦

toemoji() {
  local in=/dev/stdin
  local out=/dev/stdout
  local file="$1"
  if [[ -n $file ]];then
    in="$file"
    out="$(mktemp)"
  fi
  IFS=$'\n'
  while read -r line; do
    for k in ${!emoji[@]}; do
      line=${line//:$k:/${emoji[$k]}}
    done
    echo "$line" >> $out
  done < "$in"
  if [[ -n $file ]]; then
    cp "$out" "$file"
    rm "$out"
  fi
}

toemoji "$@"
```

Here's how you would change VIM:

```
autocmd BufWritePost *.md silent !toemoji %
```

Total all lines beginning with numbers:

```bash
#!/usr/bin/env bash

while read line; do
  n=${line%%|*}
  [[ ${line##*|} == TOTAL ]] && continue
  ((total+=n))
  echo $line
done
echo $total\|TOTAL
```

