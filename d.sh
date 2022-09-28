#!/usr/bin/env bash

function foo() {
local ret=$(cat ~/.sfdxcommands.json | jq -r ".[] | select(.id==\"$selected\") | .flags | keys[]" | $(__fzfcmd) -m --bind='ctrl-z:ignore,alt-j:preview-down,alt-k:preview-up' --preview='cat ~/.sfdxcommands.json | jq -r ".[] | select(.id==\"'$selected'\") | .flags | to_entries[] | select (.key==\""{}"\") | [\"Command:\n'"$fullcmd"'\n\",\"Flag Description:\",.value][]"' --preview-window='right:wrap')
echo "${ret//$'\n'/ --}"
}

foo
