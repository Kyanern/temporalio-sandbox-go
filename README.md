# temporalio-sandbox-go

A sandbox to experiment and play around with temporal.io and its Go SDK.

## Subprojects

* lr1
    * A simple demonstration of a branch in a workflow. Note that Temporal designed workflows to be deterministic, so Activity_randLR is an anti-example.
* lr2
    * A simple demonstration of a branch in a workflow. Unlike lr1, the path to take is decided by the client (i.e. starter/main.go).
* signals
    * A simple demonstration of using signals to determine the path of a workflow that is waiting for signals.
    * Treat this as code snippets to learn from, modify and build upon to fit more complex use cases.
