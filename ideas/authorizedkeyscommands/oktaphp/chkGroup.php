#!/usr/bin/env php
<?php
// application.php
// Okta PHP SDK a bit easier to use than the golang one

require __DIR__.'/vendor/autoload.php';

$client = (new \Okta\ClientBuilder())
            ->setConfigFileLocation('.okta/okta.yaml')
            ->build();

if (isset($argv)) {
    if (count($argv) < 3) {
        echo "Name or Groupname argument missing\n";
    } else {
        $username = $argv[1];
        $groupname = $argv[2];

        $user = new \Okta\Users\User();
        $foundUser = $user->get($username);
       // dump($foundUser);
        $groups = $foundUser->getGroups(['q' => $groupname]);
        $groupsArr = [];
        foreach ($groups as $group) {
            $groupsArr[] = $group->getProfile()->getName();
        }
        dump($groupsArr);
        dump($groupname);
       // I don't think this line works. It is simple but the if/else does
       // return in_array($groupname, $groupsArr) ? 0 : 1;
       if (in_array($groupname, $groupsArr))
        {
        echo "'$groupname' found in the array";
        }
        else
        echo "'$groupname' not found";
        ## I really only want a simple 0 or 1 exit code here

    }    
} else {
    echo "Name and GroupName arguments missing\n";
}