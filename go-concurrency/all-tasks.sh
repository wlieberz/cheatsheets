#!/bin/bash

scripts=(
task-1.sh
task-2.sh
task-3.sh
)

for script in "${scripts[@]}"
do
   : 
   ./$script 2>&1 | tee output-bash/${script}_output.json
done