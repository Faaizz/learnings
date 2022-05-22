# Microservices Security
Notes from Microservices: Security course on LinkedIn Learning.

## Content
- [Microservices Security](#microservices-security)
  - [Content](#content)
  - [IODC Tips](#iodc-tips)
  - [IODC Useful Resources](#iodc-useful-resources)
  - [Keycloak Tips](#keycloak-tips)
    - [Libraries and Adapters](#libraries-and-adapters)

## IODC Tips
- ID token should only be used for User ID on the Client and never on the API. To obtain user information, the API should query the /userinfo endpoint of the identity provider using the provided access token.
- Always include expiration dates in tokens.


## IODC Useful Resources
- [https://openidconnect.net](https://openidconnect.net)


## Keycloak Tips
### Libraries and Adapters
- Do not set realm-public-key to enable Keycloak download it automatically whenever it is needed, and prevent breaking of the library when Keycloak automatically retates its keys.
