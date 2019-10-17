#!/bin/bash

check_pipelines() {
c=$(curl https://wings.pivotal.io/api/v1/jobs 2>/dev/null | jq -r '.[] | select(.team_name=="navcon")' | jq -r .finished_build.status | grep failed | wc -l)
echo $c
}

display_sleep() {
pmset displaysleepnow
}

wake() {
NOW=$(date -v+5S +"%m/%d/%y %H:%M:%S")
pmset schedule wake "${NOW}"
}

cleanup() {
if [ -f ./stay_awake ]
then
	rm ./stay_awake
fi
}
main() {

cleanup
while [ true ]
do
	if [ $(check_pipelines) -gt 0 ]
	then
		wake
	else
		if [ -f ./stay_awake ]
		then
			wake
			exit
		else
	  display_sleep
		fi
	fi
sleep 2
done
}

main "$@"
