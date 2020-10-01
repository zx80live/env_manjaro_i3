### Manjaro i3 install

> TODO optimize GRUB - remove unused entries


##### Current version

`manjaro-i3-20.1-200913-linux58.iso`

```bash
$ screenfetch
 ██████████████████  ████████     ------
 ██████████████████  ████████     OS: Manjaro 20.1 Mikah
 ██████████████████  ████████     Kernel: x86_64 Linux 5.8.6-1-MANJARO
 ██████████████████  ████████     Uptime: 4h 17m
 ████████            ████████     Packages: 1048
 ████████  ████████  ████████     Shell: bash 5.0.18
 ████████  ████████  ████████     Resolution: 1920x1080
 ████████  ████████  ████████     WM: i3
 ████████  ████████  ████████     GTK Theme: Adapta-Nokto-Eta-Maia [GTK2/3]
 ████████  ████████  ████████     Icon Theme: Papirus-Adapta-Nokto-Maia
 ████████  ████████  ████████     Font: Noto Sans 10
 ████████  ████████  ████████     Disk: 40G / 459G (10%)
 ████████  ████████  ████████     CPU: Intel Core i7-9750H @ 12x 4.5GHz [58.0°C]
 ████████  ████████  ████████     GPU: Quadro T1000
                                  RAM: 1733MiB / 7608MiB

```



##### Download installation ISO

`https://manjaro.org/download/` 🠖  `Editions` 🠖 `Community` 🠖 `i3` 🠖 `Download`  🠖 ( `manjaro-i3-xx.xx-xxxxxx-linuxXX.iso`, `GPG Signature`,  `SHA1`)

```bash
$ mv latest.sig manjaro-i3-xx.xx-xxxxxx-linuxXX.iso.sig
```



##### Verify installation ISO
```bash
$ pacman-key -v manjaro-i3-20.1-200913-linux58.iso.sig
```

Docs: https://wiki.manjaro.org/index.php?title=How-to_verify_GPG_key_of_official_.ISO_images

```bash
# setup GPG
$ sudo pacman -S gnupg wget
$ wget github.com/manjaro/packages-core/raw/master/manjaro-keyring/manjaro.gpg
$ gpg --import manjaro.gpg
$ gpg --keyserver hkp://pool.sks-keyservers.net --search-keys 11C7F07E

# verify GPG
$ cd ~/Downloads
$ ls
manjaro-i3-20.1-200913-linux58.iso  manjaro-i3-20.1-200913-linux58.iso.sig
$ gpg --verify manjaro-i3-20.1-200913-linux58.iso.sig
$ sha1sum manjaro-i3-20.1-200913-linux58.iso
ccac883a02aad952063fc823c919ba6aeb445977              # should be equal SHA1 on site
```



##### Write ISO to USB

```bash
$ lsblk
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
sda           8:0    1  28,8G  0 disk                     # target USB flash
├─sda1        8:1    1   2,1G  0 part 
└─sda2        8:2    1     4M  0 part 
nvme0n1     259:0    0 238,5G  0 disk 
├─nvme0n1p1 259:1    0   260M  0 part 
...

# write iso image to the USB
$ sudo dd bs=4M if=~/Downloads/manjaro-i3-20.1-200913-linux58.iso of=/dev/sda status=progress oflag=sync
```



##### Initial setup screen

Select `tz=UTC`, `keyboard=en`, `drivers=nonfree` in the initial setup screen.

WARN: if you select `tz` or `keyboard` with different values, for example `keyboard=ru` then you will not be able to change keyboard layout in the console



##### Disk analyze

```bash
# determine GPT or MBR
$ sudo parted -l
```



##### Disks structure

> TODO swap with hibernate

You can change disk structure from LiveManjaro 🠖 `gparted`

WARN: there is the one attempt of configuring each partition in the partition editor. If you configure a some partition and then configure it again, then two these actions will be applied. Please verify partitions edit summary log before it will be applied.

```bash
$ lsblk
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT FORMAT FLAGS
nvme0n1     259:0    0 238,5G  0 disk                                    # SSD INTEL 250 Gb
├─nvme0n1p1 259:1    0   260M  0 part              -                     # fat32 SYSTEM EFI
├─nvme0n1p2 259:2    0    16M  0 part              - 
├─nvme0n1p3 259:3    0 237,2G  0 part /mnt/win     -                     
└─nvme0n1p4 259:4    0  1000M  0 part              -

nvme1n1     259:5    0 465.8G  0 disk                                    # SSD Samsung 500 Gb
├─nvme1n1p1 259:6    0   550M  0 part /boot/efi   [F]  [boot]            # grub
├─nvme1n1p2 259:7    0   100G  0 part /           [F]
├─nvme1n1p3 259:8    0 363.2G  0 part /home        -
└─nvme1n1p4 259:9    0     2G  0 part [SWAP]      [F]
```



#### System configuration

##### SSD TRIM setup

```bash
$ sudo systemctl status fstrim.timer
$ sudo systemctl start fstrim.timer
```



##### Swap

```bash
$ cat /proc/sys/vm/swappiness
60
$ sudo vim /etc/sysctl.d/100-manjaro.conf

vm.swappiness=10

$ sudo reboot
```



##### Synchronize date and time: manual

```bash
$ sudo ntpd -qg
```



##### Synchronize date and time: automatically

`manjaro-settings-manager` 🠖 `Time and Date` 🠖 turn on `Set time and date automatically` check-box



##### Firewall

`gufw` 🠖 enable `Status` check-box



##### Switch to the local mirror

`pamac-manager` 🠖 `Preferences` 🠖 `Official repositories` 🠖 `Use mirrors from:`🠖 select country and then press `Refresh Mirrors List`



##### Enable AUR

`pamac-manager` 🠖 `Preferences` 🠖 `AUR` 🠖 turn on `Enable AUR support` check-box and `Check for updates`



##### Update repositories

```bash
$ sudo pacman -Syyu
```



##### Base packages

```bash
$ sudo pacman -S git gvim wget mc
```



##### Keyboard layout

```bash
$ setxkbmap -layout us,ru -option 'grp:shifts_toggle'
```

Also can be defined in the config `~/.i3/config`:

```ini
exec_always --no-startup-id setxkbmap -layout us,ru -option 'grp:shifts_toggle'
```



##### Disable login attempts

```bash
# unlock locked user: CTRL+ALT+F2 and then 
$ faillock --reset --user username
# disable login attempts
$ sudo vim /etc/security/faillock.conf
deny = 0
```



##### GPG

> TODO config



##### Boot performance

> TODO configure

https://wiki.archlinux.org/index.php/Improving_performance/Boot_process



#### Drivers

##### NVIDIA

```bash
$ mhwd -l
$ sudo mhwd -a pci nonfree 0300
$ sudo reboot
$ sudo nvidia-settings
```



##### Printers

Brother DCP -L2560DWR

`pamac-manager` 🠖 `brother-dcp-l2550dw`



#### Tools: productivity

##### ncdu

`sudo pacman -S ncdu`



##### zsh

> TODO autosuggestion /h/w/ -> tab does not work

`zsh`, `zsh-autosuggestions`, `zsh-syntax-highlighting` packages are installed and enabled by default in the `.zshrc`

```bash
# change shell
$ echo $SHELL
$ chsh -l
$ chsh -s /usr/bin/zsh
$ zsh
# logout
```

```bash
# disable zsh manjaro-config
$ vim ~/.zshrc
# comment all lines
```

```bash
# enable history
$ vim ~/.zshrc
HISTFILE=~/.zsh_history
SAVEHIST=10000
```



##### zsh plugins

`pamac-manager` 🠖  `zsh-autosuggestions`, `zsh-syntax-highlighting`

```bash
$ echo 'source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.plugin.zsh' >>! ~/.zshrc
$ echo 'source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.plugin.zsh' >>! ~/.zshrc
```



##### Powerlevel10k

`pamac-manager` 🠖 `AUR: zsh-theme-powerlevel10k`, `ttf-meslo-nerd-font-powerlevel10k`

```bash
$ echo 'source /usr/share/zsh-theme-powerlevel10k/powerlevel10k.zsh-theme' >>! ~/.zshrc
```

```bash
# configuring powerlevel10k
$ p10k configure
```



##### Terminal

`pamac-manager` 🠖 `alacritty`

```bash
# configure alacritty
$ mkdir -p ~/.config/alacritty
$ cp /usr/share/doc/alacritty/example/alacritty.yml
$ vim ~/.config/alacritty/alacritty.yml
env:
  TERM: xterm-256color    # fix mc in ssh connection
...
font:
  size: 13
  normal:
    family: MesloLGS NF
    style: Regular
...
selection:
  save_to_clipboard: true # automatically copy selection to clipboard
```

```bash
# configure i3 terminal
$ vim ~/.i3/config
...
bindsym $mod+Return exec alacritty

```



##### Vim

Dependencies - `git`, `cmake`, `go`, `nodejs`, `npm` , `jdk-openjdk`

Install `gvim` package which contains already `vim` with enabled `xterm_clipboard` property. This allows working with system clipboard

```bash
$ vim --version | grep `xterm_clipboard`
... +xterm_clipboard    // system clipboard is supported
... -xterm_clipboard    // system clipboard is not supported
```

Plugin manager - https://github.com/junegunn/vim-plug

Plugin repository - https://vimawesome.com/

You should run `vim` before the next steps to download enabled plugins automatically.

```bash
# setup YouCompleteMe plugin
$ cd ~/.vim/bundle/youcompleteme
$ python3 install.py --all
```

```bash
# setup darcula theme https://github.com/blueshirts/darcula
$ mkdir ~/.vim/colors
$  wget https://raw.githubusercontent.com/blueshirts/darcula/master/colors/darcula.vim -O ~/.vim/colors/darcula.vim
```

Working with system clipboard:

```
"+p     // paste from system clipboard in NORMAL mode
ctrl+V  // select text
"+y     // copy to system clipboard
```



##### Desktop notification

```bash
$ notify-send 'Hello, world'
$ notify-send "Hello, world"
$ sudo pacman -Syyu && notify-send 'Pacman update has been completed'
```



##### Rofi

```bash
$ sudo pacman -S rofi
```

```bash
$ vim ~/.i3/config

bindsym Mod1+Tab exec rofi -show window -font "MesloLGS NF 16" -show-icons -theme sidebar
bindsym Mod1+F2 exec rofi -show run -font "MesloLGS NF 16" -show-icons -theme blue
```



##### Screenshots

> TODO upload to server
>
> TODO gpg alice

Dependencies - `xclip`, `scrot`

Use `scrot` instead of  `i3-scrot` usage in `~/.i3/config`

```
# screenshots                                                                     
set $scrdir ~/Pictures                                                            
bindsym Print            --release exec "scrot -z -e  'mv $f $scrdir && xclip -selection clipboard -t image/png -i $scrdir/$n'; sleep 1; exec notify-send 'screenshot has been saved to $scrdir'"
bindsym $mod+Print       --release exec "scrot -z -ue 'mv $f $scrdir && xclip -selection clipboard -t image/png -i $scrdir/$n'; sleep 1; exec notify-send 'screenshot has been saved to $scrdir'"
bindsym $mod+Shift+Print --release exec "scrot -z -se 'mv $f $scrdir && xclip -selection clipboard -t image/png -i $scrdir/$n'; sleep 1; exec notify-send 'screenshot has been saved to $scrdir'"

```



##### Screen recorders

> TODO upload to server
>
> TODO gpg alice

`pamac-manager` 🠖 `peek`

`pamac-manager` 🠖 `vokoscreenNG`



##### Network monitor

`pamac-manager` 🠖 `speedometer`

```
# get network interfaces
$ ip address
...
3: wlp82s0
..
$ speedometer -r wlp82s0
```



##### Time tracking

https://superuser.com/a/611582

`pamac-manager` 🠖 `termdown`



##### Backup

`grsync`



##### Bookmarks

> TODO clarify
>



##### Knowledge base

> TODO clarify
>



##### Education scheduler

> TODO clarify
>



##### Education learning log

> TODO clarify
>



#### i3 preferences

###### i3 config

Location `~/.i3/*`



###### i3 menu bar

```bash
$ cp /etc/i3status.conf ~/.i3status.conf
```



###### i3 rename window title

Dependencies `xdotool`, `wmctrl`



###### i3 save layout

Dependencies `perl-anyevent-i3`

```bash
# save layout to json
$ i3-save-tree --workspace 1 > ~/.i3/ws-layout.json
$ vim ~/.i3/ws-mount.json
# complete correct json - uncomment and define all fields

$ vim ~/.i3/config
# bind key and populate layout by the exec seqence
# each application will be placed to the proper place which is defined by regex in the `ws-layout.json`
bindsym ... exec i3-msg "append_layout ~/.i3/ws-layout.json" ; exec ...

```

Example: how to define layout cell which will contain particular running app

```bash
$ alacritty -t watch_lsblk -e watch lsblk  # run terminal with title and application 
$ xprop                                    # and then select tile which contain above executed `alacritty`
                                           # find string `WM_CLASS(STRING) = "Alacritty", "Alacritty"`
                        # find string `WM_NAME(STRING) = "watch_lsblk"`
```

Define regex in the layout template `$ vim ~/.i3/ws-layout.json`

```json
...
 "name": "watch_lsblk",                                                                  
 "percent": 0.1,
 "swallows": [
   {
     "class": "^Alacritty$",
     "instance": "^Alacritty$",
     "title": "^watch_lsblk$"           // define regex for application
   }
...
```



#### Tools: development

##### Typora

`typora`- markdown editor



##### Virtualbox

###### Virtualbox video driver

`Settings` 🠖 `Display` 🠖 `Graphics Controller` 🠖 `VboxSVGA`

https://wiki.manjaro.org/index.php?title=VirtualBox

```bash
$ mhwd-kernel -li
Currently running: 5.8.6-1-MANJARO (linux58)
$ sudo pacman -Syu virtualbox linux58-virtualbox-host-modules
$ sudo modprobe vboxguest vboxvideo vboxsf
$ sudo vboxreload
```



###### Virtualbox FTP

```bash
# install FTP on guest machine
$ sudo pacman -S vsftpd
$ sudo systemctl enable vsftpd.service
$ sudo systemctl start vsftpd.service

# config FTP on guest machine
$ sudo vim /etc/hosts.allow
vsftpd: ALL                     # Allow all connections
vsftpd: 10.0.0.0/255.255.255.0  # IP address range

$ sudo vim /etc/vsftpd.conf
...
write_enable=YES
local_enable=YES

$ sudo systemctl restart vsftpd.service
```



###### Virtualbox networking

1. Check current state on host machine

   ```bash
   $ ip address show
   ...
   2: wlp82s0                  # main wireless adapter
   inet 192.168.1.67/24        
   ...
   ```

2. Add second network adapter

   ```
   Machine 🠖 Settings 🠖 Network 🠖 Adapter 2 🠖 Bridged adapter 🠖 wlp82s0 (host main wireless adapter)
   ```

3. Check networking on guest machine

   ```bash
   $ ip address show
   ...
   2: enp0n3                   # Adapter 1 NAT (internet)
   inet 10.0.2.15/24
   ...
   3: enp0s8                   # Adapter 2 Bridged (local network)
   inet 192.168.1.48/24
   ...
   
   # ping from guest to host machine
   $ ping 192.168.1.67
   ```

4. Check networking on host machine

   ```bash
   $ ping 192.168.1.48
   $ ftp -p 192.168.1.48
   ```




