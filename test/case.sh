#!/bin/bash

case $1 in
    foo) echo "FOOOOO!" ;;
    bar) echo "BAAAAAR!" ;;
    baz) echo "BAAAZZZ!" ;;
    *) echo "WHATEVER!"
esac
