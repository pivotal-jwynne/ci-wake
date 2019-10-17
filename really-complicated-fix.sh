#!/bin/bash

fly -t wings-navcon sp -n -c pipelines/fixed.yml -p navcon-really-important-stuff-happens-here 1>/dev/null 

fly -t wings-navcon trigger-job -j navcon-really-important-stuff-happens-here/solving-all-the-worlds-problems 1>/dev/null

