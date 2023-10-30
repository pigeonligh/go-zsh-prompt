export TERM=screen
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

function _suggest {
    echo -n "suggest\0" >&4
    echo -n "$1\0" >&4
    echo -n "$2\0" >&4
    read ret <&3
    echo -n $ret
}

function __push_tab {
    if [ "$BUFFER" != "" ]; then
        NEWBUF=$(_suggest $CURSOR "$BUFFER")
        NEWCS=`expr $CURSOR + ${#NEWBUF} - ${#BUFFER}`
        BUFFER=$NEWBUF
        CURSOR=$NEWCS
        # zle end-of-line
    fi
}

zle -N __push_tab
bindkey '^I' __push_tab

export HISTFILE=$HOME/.zsh_history
export HISTSIZE=1000
export SAVEHIST=1000
setopt sharehistory
setopt appendhistory
