# HTTP Forward Proxy (wip)

This is a simple forward proxy written in golang. It can take a request and forward it to the actual requested URL, and return the actual response.

## Current status

- Forwards request after setting `X-Forwarded-For` and returns response
- Removes any well known hop-by-hop headers or any custom ones defined in `Connection`
- Bad logging

### Goals

- [x] Make sure response is proper before returning to client
- [ ] Add a website & keyword banlist
    - [ ] Web UI to manage banlists
- [ ] Handle HTTPS (need learnings)
- [ ] Make configurable and deployable (binary output)
- [ ] Add logging (hehe bye privacy)

### More

- [ ] Write a load balancer to sit in front of multiple proxies (!!)
