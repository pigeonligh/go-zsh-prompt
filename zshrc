export TERM=xterm-256color
export PS1="> "
cd $HOME

function _init {
    echo -n "init\0" >&4
    read src <&3
    source <(echo "$src")
}

_init

function _send {
    echo -n "handle\0" >&4
    echo -n "$@\0" >&4
    read src <&3
    source <(echo "$src")
}

function __push_enter {
    if [ -n "$BUFFER" ]; then
        if [ "${BUFFER: -1}" != "\\" ]; then
            print -S "$BUFFER"

            echo ""
            echo -n -e "\033[0J"
            _send "$BUFFER"
            BUFFER=""
        else
            BUFFER="${BUFFER:0:${#BUFFER}-1}"
            echo -n "\n>  "
            return
        fi
    fi
    zle accept-line
}

zle -N __push_enter
bindkey '^M' __push_enter

_suggest() {
    echo -n "suggest\0" >&4
    echo -n "$1\0" >&4
    echo -n "$2\0" >&4
    echo "suggest $1 $2" >> ./log

    while true; do
        read ret <&3
        echo "read $ret" >> ./log
        if [ -z "$ret" ]; then
            break
        fi
        echo "$ret"
    done
}

_completion() {
    local -a completions

    completions=$(_suggest $CURSOR "$BUFFER")
    completions=("${(f)completions}")

    if (( ${#completions[@]} > 0 )); then
        _describe 'custom completions' completions
    else
        compadd -- "No matches"
    fi
}

autoload -Uz compinit
compinit
for command completion in ${(kv)_comps:#-*(-|-,*)}; do
    compdef -d "$command"
    compdef -d "$completion"
done
compdef _completion -default-
setopt no_list_rows_first
setopt auto_menu
setopt auto_list

export HISTFILE=$HOME/.zsh_history
export HISTSIZE=1000
export SAVEHIST=1000
setopt sharehistory
setopt appendhistory

export PATH=
