# wad2linkedin

This is a small cli tool that helps to connect with all the great people one has met during the WeAreDevelopers conference.

Unfortunately, the export from the WAD-App does not include the LinkedIn profiles, so I build this tool to help with
connecting the people on LinkedIn.

What is does:
 - It reads the Excel export from the WAD app
 - It extracts the relevant information of the contacts
 - It opens the browser pointing to the search results on LinkedIn

More details are shown when running the CLI's help message:

```bash
wad2linkedin --help
```

## Build

You'll need to have Go installed to be able to run the following script which will create binaries for various platforms
```bash
./build.sh
```
