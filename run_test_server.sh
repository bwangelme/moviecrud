#!/usr/bin/env bash

go build ./main.go && MOVIEDEMO_ENV="TEST" ./main
