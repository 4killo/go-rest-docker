#!/bin/bash

sleep 5
if curl -X GET -H "Content-Type: text/html" -H "Cache-Control: no-cache"  service/healthcheck | grep -q 'Good!' ; then
	echo "Tests passed!"
	exit 0
else
	echo "Tests failed!"
	exit 1
fi
