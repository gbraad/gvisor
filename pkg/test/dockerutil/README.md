# dockerutil

This package is for creating and controlling docker containers for testing
runsc, gVisor's docker/kubernetes binary. A simple test may look like:

```
 func TestSuperCool(t *testing.T) {
   ctx := context.Background()
   c := dockerutil.MakeConatiner(ctx, t)
   got, err := c.Run(ctx, dockerutil.RunOpts{
     Image: "basic/alpine"
   }, "echo", "super cool")
   if err != nil {
      t.Fatalf("err was not nil: %v", err)
   }
   want := "super cool"
   if !strings.Contains(got, want){
     t.Fatalf("want: %s, got: %s", want, got)
   }
 }
```

For further examples, see many of our end to end tests elsewhere in the repo,
such as those in //test/e2e or benchmarks at //test/benchmarks.

dockerutil uses the "official" docker golang api, which is
[very powerful](https://godoc.org/github.com/docker/docker/client). dockerutil
is a thin wrapper around this API, allowing desired new use cases to be easily
implemented.

## Profiling

dockerutil is capable of generating profiles. Currently, the only option is to
use pprof profiles generated by `runsc debug`. The profiler will generate Block,
CPU, Heap, Goroutine, and Mutex profiles. To generate profiles:

*   Add the `--profile` flag to the runtime under test's config in
    `/etc/docker/daemon.json`.
*   Run the test binary doing the profiling as root. `runsc debug` inspects
    `/var/run/docker` among other things, so it needs to be root.
*   Use the `pprof` flags in dockerutil.go.
*   Create profiled containers with `dockerutil.MakeContainer()` and
    non-profiled containers with `dockerutil.MakeNative()`.

The easiest way to do this is to look at `//scripts/benchmark_profile.sh`, copy
and modify it for your purposes.

### Profiling Details

The below shows an example of using profiles with dockerutil.

```
  ctx := context.Background()
  // profiled and using runtime from dockerutil.runtime flag
  profiled := MakeContainer()

  // not profiled and using runtime runc
  native := MakeNativeContainer()

  err := profiled.Spawn(ctx, RunOpts{
    Image: "some/image",
  }, "sleep", "100000")
  // profiling has begun here
  ...
  expensive setup that I don't want to profile.
  ...
  profiled.RestartProfiles()
  // profiled activity
```