** 
These are iptables utility firewall commands 
**

**List of learned commands:

sudo apt install iptables- if suddenly the iptables utility is not installed

sudo iptables -L- display lists of current rules

sudo iptables -A INPUT -j LOG- add rule: log all incoming requests

sudo iptables -F INPUT- clear all rules in the INPUT chain

sudo iptables -A OUTPUT -j LOG- add rule: log all outgoing requests

sudo iptables -A FORWARD -j LOG- add rule: log all transit requests

hostname -I- view IP address

ip a- display information about network connections (including IP address)

sudo iptables -t filter -A INPUT -s 192.168.0.104 -j DROP- add a rule to the filter table, in the INPUT chain: reject all requests from the specified IP address

sudo iptables -t filter -A INPUT -s 192.168.0.107 -j ACCEPT- add a rule to the filter table, in the INPUT chain: allow all requests from the specified IP address

sudo iptables -L -nv --line-numbers- display a list of current rules with more detailed information

sudo iptables -t filter -D INPUT 4- delete rule number 4 in the INPUT chain

sudo iptables -A OUTPUT -p tcp --dport 80 –j REJECT- reject outgoing requests on port 80 (HTTP requests)



sudo iptables -P INPUT DROP- set the DROP policy on the incoming request rule chain

sudo iptables -P INPUT ACCEPT- allow all incoming requests by default (policy)

sudo iptables -I INPUT -p tcp --dport 22 -j ACCEPT- allow incoming requests on port 22 (SSH requests)

sudo iptables-save > ~/iptables_config- save iptables rules to a text file

sudo iptables-restore < ~/iptables_config- load iptables rules from a text file**
