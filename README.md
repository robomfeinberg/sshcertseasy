# sshcertseasy
Helping Make SSH Certificates more easy. (yes I know) 


## How can we build SSH Certificates into our DevOps processes?


```
                       +------------------+
                       |                  |
                       |  Okta            |<------+
                       |                  |       |   AuthorizedKeysCommand
                       +------------------+       |   Get list of Okta groups
                                                  |   the user is assigned
                                                  |   permits access with signed key
+-------------+                                   +---------------------+
|             |                                   |                     |
|             <----------------------------+------>                     |  user account is added on 
|    CA       |                               +---+   AWSHost           |  the fly if it does not 
|  step-ca    <---------+                     |   |                     |  exist
+-------------+         |                     |   | AuthorizedKeysCommand
                        |                     |   +---------------------+
                        |                     |
                        |                     |
                        |                     |
                        |                     |
                        |                     |
                        |                     |
                        |+--------------------|
                        ||                    +
                        +>                    |  UsesProxyCommand to generate ssh certificate
                         |  End User Laptop   |  using step ca ssh. Authenticates with SAML
                         |                    |  to Okta. Completes ssh to awshost with signed
                         | ssh awshost        |  certificate
                         +--------------------+
```                         
