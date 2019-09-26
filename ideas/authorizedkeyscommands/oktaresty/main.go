package main
// Don't look here. I have no idea what I did to muck it up so bad
//  List all members in group
//  /api/v1/groups/${groupId}/users
//  List groups member belongs to
//  /api/v1/users/${userId}/groups
// ➜  jq -r '.[] | .profile.name' json.txt
// Everyone
//HostGroup

import (
    "fmt"
	 //"encoding/json"
    "github.com/go-resty/resty"
)
 // Create a Resty Client

func main () {

//	apiToken := os.Getenv("OKTA_API_TOKEN")
apiToken := curl -v -X GET \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-H "Authorization: SSWS ${api_token}" \
"https://{yourOktaDomain}/api/v1/users/OKTASERID/groups"
client := resty.New()
// curl -v -X GET \
// -H "Accept: application/json" \
// -H "Content-Type: application/json" \
// -H "Authorization: SSWS ${api_token}" \
resp, err := client.R().//SetResult(&Authorization).
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "SSWS "+apiToken).
		Get("https://OKTADOMAIN.okta.com/api/v1/users/YOUROKTAUSERNAME/groups")
		
printOutput(resp, err)

}

// how the hell do you unmartial json with resty?

func printOutput(resp *resty.Response, err error) {
	// add error trap here?
	fmt.Println(resp, )
}
// cat json.txt | jq -r  ".[] | .profile .name"
// Explore response object
// fmt.Println("Response Info:")
// fmt.Println("Error      :", err)
// fmt.Println("Status Code:", resp.StatusCode())
// fmt.Println("Status     :", resp.Status())
// fmt.Println("Time       :", resp.Time())
// fmt.Println("Received At:", resp.ReceivedAt())
//fmt.Println("Body       :\n", resp)
//fmt.Println()

//fmt.Println("Token		:",apiToken)
// Explore trace info
// fmt.Println("Request Trace Info:")
// ti := resp.Request.TraceInfo()
// fmt.Println("DNSLookup    :", ti.DNSLookup)
// fmt.Println("ConnTime     :", ti.ConnTime)
// fmt.Println("TLSHandshake :", ti.TLSHandshake)
// fmt.Println("ServerTime   :", ti.ServerTime)
// fmt.Println("ResponseTime :", ti.ResponseTime)
// fmt.Println("TotalTime    :", ti.TotalTime)
// fmt.Println("IsConnReused :", ti.IsConnReused)
// fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
// fmt.Println("ConnIdleTime :", ti.ConnIdleTime)