go-cups-control-files
========

# What's this?
It's rather difficult to extract more than the common metadatas from a printing job in cups. Cups generates so called control files which contain a lot more metadatas than the command line tools or the gui can provide. We were especially interested in the start and completion time. On stackoverflow Kurt Pfeifle(https://stackoverflow.com/questions/53688075/how-to-dissect-a-cups-job-control-file-var-spool-cups-cnnnnnn) found out, that there is a tool for reading in these control files.

This tool is a go clone of the testipp tool. Because I was unable to compile cups staticly and I would like to be able later to use it in other go based applications I've rewritten it on a weekend in go :)

# What it can (currently) do
It can currently dump nearly all attributes of a control file in json. The attribute job-sheets is currently not implemented yet. I haven't figured out the enum format currently :)

# How to use
First, you have to clone and build it - no releases are currently provieded

```
go build
```

For analyzing a control file, just run:
```
./go-cups-control-files /tmp/c00325
```

For nicer formatting pipe it to jq
