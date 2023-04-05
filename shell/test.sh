#!/bin/bash

go build -o ./dist/$DIR/ ./$DIR/$FILE.go
oj t -c ./dist/$DIR/$FILE