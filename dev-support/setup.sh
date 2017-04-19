# /bin/sh


target_file="$HOME/.bash_profile"
read -p "Warning! This script should only be executed once. Are you sure to continue?[y/n]" response
if [[ "$response" =~ ^[yY]+$ ]]
then
 	if [ -z $1 ]
	then
		echo "must input the directory of the GO workshop"
	else
		GOPATH=$1
		echo "# workspace's bin subdirectory to your GO PATH" >> $target_file
		echo "export GOPATH=$1" >> $target_file
		echo "export PATH=\$PATH:\$GOPATH/bin" >> $target_file
		source $HOME/.bash_profile
	fi
fi

 
