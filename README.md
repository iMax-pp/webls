# webls
`webls` list all files/folders under given repository and send it back as a JSON array.
The files to return can be filter by extension (no regex, default to all).

```bash
$ ./webls -help
Usage of ./webls:
  -addr string
    http service address (default ":8080")
  -ext string
    file extension to filter (default ".*")
  -repo string
    repository to expose (default ".")
```

2 endpoints:
- `files/` recursively list all files under given folder (`-repo` flag) and of given extension (`-ext` flag)
- `folders/` recursively list all folders under given folder (`-repo` flag)