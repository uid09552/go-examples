sudo apt update   
sudo apt upgrade -y    
sudo apt install nfs-kernel-server -y   
sudo mkdir /var/nfs/general -p   
sudo chown nobody:nogroup /var/nfs/general    
Edit the file /etc/exports such that the following line is added:    
/var/nfs/general      ip.address.of.node1(rw,sync,no_subtree_check) ip.address.of.nodeX(rw,sync,no_subtree_check)    
