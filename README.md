> This repository is archived and no longer being maintained.

# Must

AUR helper with apt-like syntax

## Usage

### Install

`must install <packages to install>`

Uses git to clone the AUR repo for each package and runs `makepkg` against each PKGBUILD.

### Remove

`must remove <packages to remove>`

Runs `pacman -Rns` for each provided package and removes any leftover AUR files from the filesystem.

### Update

`must update`

Uses git to download new PKGBUILDs from the AUR.

### Upgrade

`must upgrade`

Runs `makepkg -si` against packages which are known to have an upgrade available. The upgrade list can be updated using `must update`.

## Installing

To install Must, first clone the repository.

```shell
git clone https://github.com/BrenekH/must.git
```

Then change directory into the project folder.

```shell
cd must
```

Finally install using makepkg.

```shell
makepkg -sci
```

> Note: If you're unfamiliar with makepkg, "-sci" will install any required dependencies, clean up after itself and install the package when it is finished packaging using pacman -U.

## Why did I even make this?

Before building Must, I only used `git`, `makepkg`, and `pacman` to manage any software I used from the AUR.

I was fairly happy with my setup except for one pain point, updates.
Changing directory into each repository and running `git pull` was a very tedious process, especially when there were lots of updates to complete.

I first considered using a popular AUR helper such as `yay` but I decided that I really don't like the way `pacman` uses single-letter options as commands. I mean seriously, who thought that `-Syu` and `-Syyu` should do different things?[*](#pacman-pref-note)

Plus, I'm one of those crazy guys who likes to create their own tools rather than relying on someone else to build them for them.

## License

This project is licensed under the GNU Public License version 3, a copy of which can found in the [LICENSE file](https://github.com/BrenekH/must/tree/master/LICENSE).

<!-- markdownlint-disable-next-line -->
<a name="pacman-pref-note"></a>\* Pacman is a very powerful and versatile tool, thanks in part to the single-letter commands.
However, everyone has a right to like what they want, so while I prefer the subcommand structure of apt, I can't fault anyone for preferring the Pacman syntax.
This project is just not for them.
