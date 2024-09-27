# Golang Imports Practice

## How to run:

```bash
go run .
```

or 

```bash
go run main.go adjacent.go
```

## Key learnings

- Local imports are of the structure `{module}/{package directory}`
- All Go source files in the same directory must belong to the same package (one package per directory)

### Additional learnings

- Functions for use on import should be capitalized ("Exported")
- Package name doesn't have to match directory name
- Package filenames don't really matter (see `main.go` in root and in subsubfolder)
- Name functions relative to the package they're defined in
    - Instead of `PrintFromSubfolder` just name it `Print` because it's going to be called as `subfolder.Print` anyway. [[docs](https://go.dev/wiki/CodeReviewComments#package-names)]
- Naming conventions
    - camelCase for variables and function names
    - snake_case for file names
- Importing submodules (modules contained inside another module) is done by replacing the module reference with the path to the submodule. See [go.mod](./go.mod).

## Questions yet to be answered

- Is there a naming convention for module names?
    - They have to be globally unique, which is easy to do if using github to host (e.g. github.com/username/package-repo).
    - At a glance, it looks like the repo names tend to be kebab-case.
    - I don't think this actually matters though: you import the module, but it's the package name that gets used as the reference.

## Additional reading

https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/
https://go.dev/wiki/CodeReviewComments
https://go.dev/doc/effective_go
