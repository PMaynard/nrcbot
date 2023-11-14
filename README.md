# Nameless Research Collective (NRC) Bot 

A bot used to help make life easier. 

**Install**

[Golang](https://go.dev/learn/) is required to compile and run this tool. Final versions will have release binaries.

	go install github.com/PMaynard/nrcbot@latest

**Usage**

Connects to an IRC channel and listens for `!doi <url|DOI>` from a whitelisted user. It will search Crossref for the DOI then create an issue on a GitHub repository with the document title and abstract.

	nrcbot 

Expects a `config.json` file to be located in the same directory.
