# Burrito 🌯

## Description
```burrito``` aims to be a full-featured way to order Chipotle from your command line, utilizing Chipotle’s online ordering API. Even though online ordering can be easily done through a browser, ```burrito``` will enable users to utilize scripting (ordering when a particular action occurs), or can potentially serve as an adapter for other applications wishing to kickstart their project with the Chipotle API.

## Security and Risks
```burrito``` aims to only act as an extension for already established APIs (in this case Chipotle's)... In practice, using ```burrito``` should be no less safe than ordering Chipotle using the browser. Connections are established over HTTPS using TLS, and your password is not stored locally. 

```burrito``` does save a API token (basically a cookie) to ```~\.netrc``` which could be used maliciously to interact with the API on your behalf. In a practical sense, storing it here poses minimal risk.

- ```burrito``` is at the mercy of Chipotle’s API — while reading through their Terms of Service, this does seem legal (we’re not impacting their services in any way), however their API might change, causing an eternal game of cat-and-mouse… Them changing their API, and us having to refactor ```burrito``` to make it work.
- ```burrito``` won’t handle credit card information directly — we’ll prompt the user to add a card using the website first. This offloads any potential PCI compliance issues off our application.
- Chipotle rate limits their API, preventing accidental errors which might leads to 100s of orders being made, but we're considering implementing rate limiting and warnings on the client, resolving concerns of accidental overusage.
