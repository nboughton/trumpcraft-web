#!/bin/bash
cd $(readlink -f $(dirname $0))
cat lovecraft.txt trump.txt > trumpcraft.txt
cat peter_rabbit.txt lovecraft.txt > peter_rabbitcraft.txt
cat peter_rabbit.txt lovecraft.txt trump.txt > all.txt