# Open

A tool for opening URL's using their OS defined default application.
This is an abstraction around standard OS specific default program opening commands:

* Linux: [`xdg-open <URL>`](https://linux.die.net/man/1/xdg-open)
   * WSL: uses [`wslpath`](https://github.com/microsoft/WSL/issues/2715) to convert to windows path then uses the Windows opener
* Mac: [`open <URL>`](https://scriptingosx.com/2017/02/the-macos-open-command/)
* Windows: [`powershell.exe -c "Start-Process -FilePath <URL>"`](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.management/start-process?view=powershell-7.5)

After [installation](#installation), you can simply _open_ urls:

~~~bash
# open a web browser at a specific url
open https://github.com/pastdev/open

# open an image viewer
open ~/pictures/me.jpg

# open a text editor
open /tmp/app.out
~~~

## Installation

`open` is a self contained binary that has [pre-built releases for various platforms](https://github.com/pastdev/open/releases).
You may find this script valuable for installation:

~~~bash
# note this command uses clconf which can be found here:
#   https://github.com/pastdev/clconf
(
  # where do you want this installed?
  binary="${HOME}/.local/bin/open"
  # one of linux, darwin, windows
  platform="linux"
  curl \
    --location \
    --output "${binary}" \
    "$(
      curl --silent https://api.github.com/repos/pastdev/open/releases/latest \
        | clconf \
          --pipe \
          jsonpath "$..assets[*][?(@.name =~ /askai-${platform/windows/windows.exe}/)].browser_download_url" \
          --first)"
  chmod 0755 "${binary}"
)
~~~

## Development

Use the default `go run`:

~~~bash
go run ./cmd/open "https://github.com/pastdev/open"
~~~
