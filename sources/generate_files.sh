#!/bin/bash
cd $(readlink -f $(dirname $0))
cat trump.txt lovecraft.txt > trumpcraft.txt