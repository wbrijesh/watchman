# Watchman
_I am building watchman for when you need a cron job to be triggered not every predefined interval but on changes in files or subdirectories of a directory._

### How to install watchman?

To install watchman you need go installed on your computer

```sh
git clone git@brijesh.dev:watchman.git
cd watchman
go build .
sudo cp watchman /usr/bin
```

If you don't have go installed, then install from [official website](https://go.dev/dl/)

### Todo

- [x] Take CLI inputs
- [x] Create a recursive list of subdirectories and files
- [x] Watch for changes to files or subdirectories
- [ ] Watch for changes in file contents
- [x] Run command when change is detected
