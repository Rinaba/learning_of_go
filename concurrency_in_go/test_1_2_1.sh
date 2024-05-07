#!/bin/sh
CNT=1

while [ 0 ]
do
	echo $CNT
	go run 1_2_1.go
	((CNT++))
done
