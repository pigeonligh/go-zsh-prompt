export TERM=screen
export PS1="> "
cd $HOME

function _send {
    echo -n "$@\0" >&4
    read <&3
}

function __push_enter {
    if [ "$BUFFER" != "" ]; then
        print -S "$BUFFER"

        echo ""
        _send "$BUFFER"
        BUFFER=""
    fi
    zle accept-line
}

zle -N __push_enter
bindkey '^M' __push_enter

export HISTFILE=$HOME/.zsh_history
export HISTSIZE=1000
export SAVEHIST=1000
setopt sharehistory
setopt appendhistory
