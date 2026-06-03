List of learned commands:

sudo apt update- update the list of packages from repositories (preparing an update or installing software)

sudo apt upgrade -y- perform a software update

sudo apt-get install htop- install the htop utility

sudo apt-get remove htop- remove htop utility

sudo apt-get purge htop- remove the htop utility along with its settings

apt moo- cow superpower

aptitude moo -vvvvv- cow superpower 2

sudo apt install cowsay- cow superpower 3



dpkg -l- display a list of installed packages

dpkg -l | wc -l- display the number of installed packages

dpkg –l tree- display info about the tree package

apt show tree- display info about the tree package

apt download htop- download the package as a *.deb file

sudo dpkg -i htop_3.0-11_amd64.deb- install the package using the dpkg utility

sudo apt install snapd- install snap package manager

apt search opera- search for opera package in apt repositories

snap find opera- search for the opera package in the snap repositories

snap list- list packages installed by snap manager

sudo snap install opera- install opera package

sudo snap remove opera- remove opera package



# download VirtualBox deb package from the official website

wget https://download.virtualbox.org/virtualbox/7.1.4/virtualbox-7.1_7.1.4-165100~Ubuntu~jammy_amd64.deb



ls

# install deb package from downloaded file

sudo dpkg -i virtualbox-7.1_7.1.4-165100~Ubuntu~jammy_amd64.deb

# remove VirtualBox package

sudo dpkg -r virtualbox-7.1



# working with the /etc/apt directory

ls /etc/apt

cd /etc/apt

ll



# output the specified file to the screen

cat sources.list

# create a new file to specify the VirtualBox repository

sudo nano sources.list.d/vbox.list

# view catalog

ls sources.list.d

tree



# output the created file to the screen

cat sources.list.d/vbox.list

# download and install the VirtualBox repository key

wget -O- https://www.virtualbox.org/download/oracle_vbox_2016.asc | sudo gpg --yes --output /usr/share/keyrings/oracle-virtualbox-2016.gpg --dearmor

# refresh repository lists (important after adding a new repo)

sudo apt update

# install VirtualBox package from repository

sudo apt install virtualbox



# call VirtualBox help (to confirm successful installation)

whatis virtualbox

virtualbox --help



# output distribution code name

lsb_release -a



# call full help for add-apt-repository command

man add-apt-repository

# install additional package for add-apt-repository utility

sudo apt install software-properties-common
