#!/usr/bin/env bash
# very simple and dumb but working example of bash. This works! It took me all of 5 minutes.
# getGroups.sh oktusername oktagroup
USER=$1
api_token="YOURAPITOKEN"
DOMAIN="OKTADOMAIN.okta.com"
GROUP=$2

#function get_groups () {
     
curl -s -X GET \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-H "Authorization: SSWS ${api_token}" \
"https://$DOMAIN/api/v1/users/$USER/groups" #  | jq -r  ".[] | .profile .name"
    
exit 0

