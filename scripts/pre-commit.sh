#!/bin/bash
#------------------------
#
# Objective: Execute steps before commit
#
# Requisites: 
# 1- 
#
# References:
# https://hackernoon.com/automate-your-workflow-with-git-hooks-fef5d9b2a58c
# https://sigmoidal.io/automatic-code-quality-checks-with-git-hooks/
# https://stackoverflow.com/questions/4774054/reliable-way-for-a-bash-script-to-get-the-full-path-to-itself
#------------------------



#------------------------
# Variables

DEBUG=false
SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

#------------------------

# If any command inside script returns error, exit and return that error 
set -e

# Access path of script
cd $SCRIPTPATH

# Access path dir parent (consider DIR_PARENT./git/hook)
cd ../../ > /dev/null 2>&1

# Show dir current
if $DEBUG ; then
    echo "[DEBUG] Dir before of script."
    pwd
fi

gofmt -w src/
