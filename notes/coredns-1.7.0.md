+++
title = "CoreDNS-1.7.0 Release"
description = "CoreDNS-1.7.0 Release Notes."
tags = ["Release", "1.7.0", "Notes"]
release = "1.7.0"
date = 2020-03-24T10:00:00+00:00
author = "coredns"
+++

The CoreDNS team has released
[CoreDNS-1.7.0](https://github.com/coredns/coredns/releases/tag/v1.7.0).

This is a **backwards incompatible release**. Major changes include:
* Better [metrics names](https://github.com/coredns/coredns/pull/3776).
* New `transfer` plugin that removes the need for plugins to perform their own zone transfers.
* The *federation* plugin (allows for v1 Kubernetes federation) has been removed. We've also removed
  some supporting code from the *kubernetes* plugin, so it will not build as an external plugin
  (with this version of CoreDNS).

As this was already backwards incompatible release, we took the liberty to stuff as much of it in
one release as possible to minimize the disruption going forward.

A new plugin, [*dns64*](https://coredns.io/plugins/dns64) as promoted from external to a plugin that
is included by default. This plugin "enables DNS64 IPv6 transition mechanism."

### Metric Changes

It's mostly dropping `count` from `_total` metrics names:

* `coredns_request_block_count_total` -\> `coredns_dns_blocked_requests_total`
* `coredns_request_allow_count_total` -\> `coredns_dns_allowed_requests_total`

* `coredns_dns_acl_request_block_count_total` -\> `coredns_acl_blocked_requests_total`
* `coredns_dns_acl_request_allow_count_total` -\> `coredns_acl_allowed_requests_total`

* `coredns_autopath_success_count_total` -\> `coredns_autopath_success_total`

* `coredns_forward_request_count_total` -\> `coredns_forward_requests_total`
* `coredns_forward_response_rcode_count_total` -\> `coredns_forward_responses_total`
* `coredns_forward_healthcheck_failure_count_total` -\> `coredns_forward_healthcheck_failures_total`
* `coredns_forward_healthcheck_broken_count_total` -\> `coredns_forward_healthcheck_broken_total`
* `coredns_forward_max_concurrent_reject_count_total` -\> `coredns_forward_max_concurrent_rejects_total`

* `coredns_grpc_request_count_total` -\> `coredns_grpc_requests_total`
* `coredns_grpc_response_rcode_count_total` -\> `coredns_grpc_responses_total`

* `coredns_panic_count_total` -\> `coredns_panics_total`
* `coredns_dns_request_count_total` -\> `coredns_dns_requests_total`
* `coredns_dns_request_do_count_total` -\> `coredns_dns_do_requests_total`
* `coredns_dns_response_rcode_count_total` -\> `coredns_dns_responses_total`

* `coredns_reload_failed_count_total` -\> `coredns_reload_failed_total`

* `coredns_cache_size` -\> `coredns_cache_entries`

And note that
`coredns_dns_request_type_count_total` is now part of `coredns_dns_requests_total` .

## Brought to You By
## Noteworthy Changes
