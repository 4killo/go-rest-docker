#!/bin/bash
name=ci
docker-compose -f docker-compose.test.yml -p $name build
docker-compose -f docker-compose.test.yml -p $name up -d
#docker wait ci_test_1

if [ $(docker wait ${name}_test_1) -eq 0 ]
then
  echo $?
  echo "Test Passed"
  docker-compose -f docker-compose.yml -p 4killo/gorest:$GIT_COMMIT build
  exit 0
else
  echo "Test Failed"
  exit 1
fi
