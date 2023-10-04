#!/bin/bash

# Mengambil argumen dari baris perintah
folder_name="$1"
facebook_username="$2"
linkedin_username="$3"

# Membuat struktur folder
mkdir -p "$folder_name"
cd "$folder_name"
mkdir -p about_me/personal about_me/professional my_friends my_system_info

# Membuat file facebook.txt dan linkedin.txt
echo "https://facebook.com/$facebook_username" > about_me/personal/facebook.txt
echo "https://linkedin.com/in/$linkedin_username" > about_me/professional/linkedin.txt

# Mengambil daftar teman dari file eksternal
curl -o my_friends/list_of_my_friends.txt "https://example.com/friends.txt"

# Membuat file about_this_laptop.txt
user_info="$(whoami)"
system_info="$(uname -a)"
echo -e "User Info:\n$user_info\nSystem Info:\n$system_info" > my_system_info/about_this_laptop.txt

# Membuat file internet_connection.txt
ping -c 3 google.com > my_system_info/internet_connection.txt

# Menampilkan struktur folder yang telah dibuat
tree "$folder_name"

# Run with  "./automate.sh "tegar at $(date +"%a %b %-d %H:%M:%S %Z %Y")" tegarimansyahfb tegarimansyahlinkedin"