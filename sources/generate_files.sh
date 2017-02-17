#!/bin/bash
cd $(readlink -f $(dirname $0))
cat lovecraft.txt trump.txt > trumpcraft.txt