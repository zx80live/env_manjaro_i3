# Archlinux install



### Prepare installation ISO

###### Download installation ISO

`https://www.archlinux.org/download/` 🠖 `torrent link`, `PGP signature (archlinux-2020.09.01-x86_64.iso.sig)`



###### Verify installation ISO

```bash
$ gpg --keyserver-options auto-key-retrieve --verify archlinux-2020.09.01-x86_64.iso.sig
$ pacman-key -v archlinux-2020.09.01-x86_64.iso.sig
```



###### Write ISO to USB

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
$ sudo dd bs=4M if=~/Downloads/archlinux-2020.09.01-x86_64.iso of=/dev/sda status=progress oflag=sync
```



### Install

##### Installer font

```bash
$ setfont /usr/share/kbd/consolefonts/LatGrkCyr-12x22.psfu.gz
or
$ setfont /usr/share/kbd/consolefonts/ter-v22n.psf.gz
```



##### wifi

```bash
$ ip address
...find network adapter
$ wpa_passphrase <network-name> >> /etc/wpa_supplicant/test.conf
...enter passphrase
$ wpa_supplicant -i <network-adapter> -c /etc/wpa_supplicant/test.conf -B
$ dhcpcd <network-adapter>
$ timedatectl set-ntp true
```



##### set large font

> TODO verify if is need

```bash
$ pacman -S terminus-font vim
$ setfont /usr/share/kbd/consolefonts/ter-v28n.psfu.gz
$ sudo vim /etc/vconsole.conf
FONT=ter-v28n
```



##### partitions
```bash
$ cfdisk
$ fdisk -l
```



  ##### disk with Windows 10
```bash
/dev/nvme0n1p1  260M EFI
/dev/nvme0n1p1  Microsoft data
...
```



  ##### create disk for archlinux

> TODO clarify disk structure

```bash
/dev/nvme1n1p1  512M EFI
/dev/nvme1n1p2  2G   swap
/dev/nvme1n1p3  100G /
/dev/nvme1n1p4  ...G /home
```

```bash
$ mkfs.fat -F32 /dev/nvme1n1p1  
$ mkswap /dev/nvme1n1p3
$ swapon /dev/nvme1n1p3  
$ mkfs.ext4 /dev/nvme1n1p3
$ mkfs.ext4 /dev/nvme1n1p4
```

```bash
$ mount /dev/nvme1n1p3 /mnt
$ mkdir /mnt/boot
$ mount /dev/nvme1n1p1 /mnt/boot
$ mkdir /mnt/home
$ mount /dev/nvme1n1p4 /mnt/home
```



##### Install
```bash
$ pacstrap /mnt base linux linux-firmware
$ genfstab -U /mnt >> /mnt/etc/fstab
$ arch-chroot /mnt
```

```bash
$ pacman -Syu
$ pacman -S mc wpa_supplicant dhcpcd net-tools vim iproute2 dosfstools mtools sudo
$ pacman -S grub efibootmgr
$ pacman -S os-prober
$ os-prober
```

```bash
$ ln -sf /usr/share/zoneinfo/Europe/Moscow /etc/localtime
$ hwclock --systohc
$ vim /etc/locale.gen
$ en_US.UTF-8
$ ru_RU.UTF-8
```

```bash
$ locale-gen
$ vim /etc/locale.conf
$ LANG=ru_RU.UTF-8
```

```bash
$ vim /etc/hostname
p53
$ vim /etc/hosts
127.0.0.1	localhost
::1		localhost
127.0.1.1	p53.localdomain	p53
```

```bash
$ passwd
$ useradd -g users -G power,storage,wheel -m <username>
$ passwd  <username>
vim /etc/sudoers
...uncomment %wheel%... line
```



##### boot loader
```bash
$ grub-install --target=x86_64-efi --efi-directory=/boot/ --bootloader-id=GRUB
$ mkdir /mnt2 # for EFI windows disk
$ mount /dev/nvme0n1p1 /mnt2
$ grub-mkconfig -o /boot/grub/grub.cfg  # should defined Windows boot manager on /dev/nvme0n1p1 EFI
```

```bash
$ exit
$ reboot
```




### Setup environment

##### xorg
```bash
$ pacman -S xorg-server xorg-xinit
$ pacman -S xorg-twm xorg-xclock xterm
```




##### i3
```bash
$ sudo pacman -S i3
...select i3-gaps i3blocks i3lock i3status
```




##### wallpaper
```bash
# google -> Kde Plasma nebula 325
$ pacman -S feh
$ vim ~/.zshrc
feh --bg-scale ~/temp/img/wallpaper.png
```



##### links

- https://www.youtube.com/watch?v=C3D_qzw94v8
- https://www.youtube.com/watch?v=jW4GFGOIUjc
- https://www.youtube.com/watch?v=gIqV4Qn2qUE
- https://erikdubois.be/changing-font-type-i3/
- https://platypus-boats.readthedocs.io/en/stable/source/Arch.html

##### mhwd -h
```bash
$ mhwd -l
> 0000:01:00.0 (0300:10de:1fb9) Display controller nVidia Corporation:
--------------------------------------------------------------------------------
                  NAME               VERSION          FREEDRIVER           TYPE
--------------------------------------------------------------------------------
video-hybrid-intel-nvidia-450xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-440xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-435xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-430xx-bumblebee            2019.10.25               false            PCI
video-hybrid-intel-nvidia-418xx-bumblebee            2019.10.25               false            PCI
    video-nvidia-450xx            2019.10.25               false            PCI
    video-nvidia-440xx            2019.10.25               false            PCI
    video-nvidia-435xx            2019.10.25               false            PCI
    video-nvidia-430xx            2019.10.25               false            PCI
    video-nvidia-418xx            2019.10.25               false            PCI
           video-linux            2018.05.04                true            PCI
     video-modesetting            2020.01.13                true            PCI
            video-vesa            2017.03.12                true            PCI


> 0000:00:02.0 (0300:8086:3e9b) Display controller Intel Corporation:
--------------------------------------------------------------------------------
                  NAME               VERSION          FREEDRIVER           TYPE
--------------------------------------------------------------------------------
video-hybrid-intel-nvidia-450xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-440xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-435xx-prime            2019.10.25               false            PCI
video-hybrid-intel-nvidia-430xx-bumblebee            2019.10.25               false            PCI
video-hybrid-intel-nvidia-418xx-bumblebee            2019.10.25               false            PCI
           video-linux            2018.05.04                true            PCI
     video-modesetting            2020.01.13                true            PCI
            video-vesa            2017.03.12                true            PCI

```

