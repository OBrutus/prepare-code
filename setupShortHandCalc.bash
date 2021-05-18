echo "This file inputs command and sets it to your path variable"
echo "This would be the command required when you will write on time"
read -p "So enter command :: " -r cmd
echo "$cmd, good choice of command."
echo "remember it this will be your command"
echo "setting up ..."

cp shortHandCalc.bash /usr/bin/$cmd

if [[ $? -ne 0 ]]
then
	echo "Opps! somethings wrong"
	echo "Might be permission issue"
	echo "Try running using sudo"
	exit 
fi

chmod +x /usr/bin/$cmd

echo "Done !"
echo "rm /usr/bin/$cmd" > uninstallShortHandCalc.sh
echo "Created uninstallation file as well"

echo "Now press Alt+F2 and use your command"
echo "Example :: Alt+F2 -> $cmd 2+20+200" 
echo "The answer should be 222"
