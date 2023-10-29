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

function _suggest {
    echo -n "$1\0" >&6
    echo -n "$2\0" >&6
    read ret <&5
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
