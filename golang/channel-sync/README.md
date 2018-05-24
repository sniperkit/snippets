# channel-sync

Implements a search algorithm for "idle" resources. Acquiring resources which are "busy" block forever.
We use the golang "context" package and channels to search in a non-blocking way.
