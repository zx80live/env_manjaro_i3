# Open a menu to find Wi-Fi connections.
# This creates a profile and connects to it
sudo wifi-menu

# Auto connect on boot
# (Assuming 'wlp3s0' is your interface name)
sudo systemctl enable netctl-auto@wlp3s0.service

# Set it to auto-connect on boot
sudo netctl-auto enable wlp3s0-WifiNameHere

# To manuall yconnect to a wifi network:
sudo netctl-auto switch-to wlp3s0-WifiNameHere
