#!/bin/bash
set -e
curl --proxy "http://localhost:4041" -v $1
